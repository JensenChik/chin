# coding=utf-8
import os
import signal
import time
from datetime import datetime
from subprocess import Popen
from model import Task, TaskInstance, DBSession, Machine
import ConfigParser
import logging
import traceback
import psutil


class Shell:
    # 具体的任务线程
    def __init__(self, command, task_id, version, log_path):
        self.task_id = task_id
        self.version = version
        now = datetime.now().strftime('%Y%m%d%H%M%S')
        self.log_path = os.path.join(log_path, '{}_{}_{}.log'.format(task_id, version, now))
        self.log = open(self.log_path, 'w')
        self.process = Popen(command, shell=True, env=dict(os.environ.items()),
                             preexec_fn=os.setsid, stdout=self.log, stderr=self.log)

    def is_running(self):
        return self.process.poll() is None

    def kill(self):
        os.killpg(os.getpgid(self.process.pid), signal.SIGTERM)

    def success(self):
        return self.process.returncode == 0

    def get_log(self):
        with open(self.log_path) as log:
            log_content = log.read()
        return log_content


class TaskTracker:
    # 负责执行任务
    def __init__(self):
        cf = ConfigParser.ConfigParser()
        cf.read('chin.ini')
        self.heartbeat_sec = int(cf.get('worker', 'heartbeat_sec'))
        self.name = cf.get('worker', 'name')
        self.running = []

        # 配置logger
        self.log_path = cf.get('worker', 'log_path')
        self.logger = logging.getLogger('worker')
        handler = logging.FileHandler(os.path.join(self.log_path, 'worker.log'))
        handler.setLevel(logging.INFO)
        handler.setFormatter(logging.Formatter('[%(levelname)s]\t%(asctime)s\t%(message)s'))
        self.logger.addHandler(handler)

        # 清理worker重启前遗留的任务
        session = DBSession()
        remain_task = session.query(TaskInstance) \
            .filter_by(execute_machine=self.name) \
            .filter_by(status='running') \
            .all()
        for instance in remain_task:
            instance.log = '由于 worker 重启，宕机前 running 的任务无法确定是否执行成功，请手动校验'
            instance.status = 'failed'
            instance.finish_time = datetime.now()
        session.commit()
        session.close()

    # 扫描数据库看是否有属于自己的任务
    def execute(self, session):
        waiting_task = session.query(TaskInstance) \
            .filter_by(execute_machine=self.name) \
            .filter_by(status='waiting') \
            .all()
        for taskInstance in waiting_task:
            task = session.query(Task).filter_by(id=taskInstance.task_id).first()
            shell = Shell(task.command, task.id, taskInstance.version, self.log_path)
            self.running.append(shell)
            taskInstance.status = 'running'
            taskInstance.begin_time = datetime.now()
            taskInstance.run_count += 1
        session.commit()

    # 杀死任务
    def kill(self, session):
        killing_tasks = session.query(TaskInstance) \
            .filter_by(execute_machine=self.name) \
            .filter_by(status='killing').all()
        for task_instance in killing_tasks:
            not_in_running_list = True

            # 常规的 killing 任务
            for running_task in self.running:
                if running_task.task_id == task_instance.task_id and running_task.version == task_instance.version:
                    running_task.kill()
                    not_in_running_list = False

            # 在worker挂掉期间killing任务
            if not_in_running_list:
                task_instance.status = False
                task_instance.log = 'worker宕机期间执行killing操作，任务不确定是否killing成功，请手动确认'
                task_instance.finish_time = datetime.now()
        session.commit()

    # 追踪任务执行
    def track(self, session):
        for shell in self.running:
            if shell.is_running(): continue
            task_instance = session.query(TaskInstance) \
                .filter_by(task_id=shell.task_id) \
                .filter_by(version=shell.version) \
                .first()
            if shell.success():
                task_instance.status = 'success'
            else:
                task_instance.status = 'failed'
            task_instance.finish_time = datetime.now()
            task_instance.log = shell.get_log()
            self.running.remove(shell)
        session.commit()

    # 反馈自身负载情况
    def health_feedback(self, session):
        men_load = psutil.virtual_memory().percent
        cpu_load = psutil.cpu_percent(3)
        machine = session.query(Machine).filter_by(name=self.name).first()
        machine.men_load = men_load
        machine.cpu_load = cpu_load
        machine.update_time = datetime.now()
        session.add(machine)
        session.commit()

    def serve(self):
        while True:
            try:
                session = DBSession()
                self.execute(session)
                self.kill(session)
                self.track(session)
                self.health_feedback(session)
                session.close()
            except Exception, e:
                self.logger.error(e)
                self.logger.error(traceback.format_exc())
            time.sleep(self.heartbeat_sec)
