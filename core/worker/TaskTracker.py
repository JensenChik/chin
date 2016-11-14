# coding=utf-8
import os
import signal
import time
from datetime import datetime
from subprocess import Popen
from model import Task, TaskInstance, DBSession
import ConfigParser


class Shell:
    # 具体的任务线程
    def __init__(self, command, task_id, version, log_path):
        self.task_id = task_id
        self.version = version
        now = datetime.now().strftime('%Y%m%d%H%M%S')
        self.log_path = os.path.join(log_path, '{}_{}_{}.log'.format(task_id, version, now))
        self.log = open(self.log_path, 'w')
        self.process = Popen(command, shell=True, preexec_fn=os.setsid, stdout=self.log, stderr=self.log)

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
        self.log_path = cf.get('worker', 'log_path')
        self.running = []

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
        pass

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
                task_instance.finish_time = datetime.now()
            else:
                task_instance.status = 'failed'
            task_instance.log = shell.get_log()
        session.commit()

    # 反馈自身负载情况
    def health_feedback(self):
        pass

    def serve(self):
        while True:
            session = DBSession()
            self.execute(session)
            time.sleep(self.heartbeat_sec)
            self.kill(session)
            self.track(session)
            session.close()
