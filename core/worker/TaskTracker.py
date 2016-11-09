# coding=utf-8

from subprocess import Popen
from datetime import datetime
import os
import signal
from datetime import datetime
from core import DBSession
from core.model import Task, TaskInstance
import time


class Shell:
    # 具体的任务线程
    def __init__(self, command, task_id, version):
        self.command = command
        self.task_id = task_id
        self.version = version
        now = datetime.now().strftime('%Y%m%d%H%M%S')
        self.log = open('{}_{}_{}.log'.format(self.task_id, self.version, now), 'w')

    def run(self):
        self.process = Popen(self.command, shell=True, preexec_fn=os.setsid, stdout=self.log, stderr=self.log)

    def is_running(self):
        return self.process.poll() is None

    def kill(self):
        os.killpg(os.getpgid(self.process.pid), signal.SIGTERM)

    def success(self):
        return self.process.returncode


class TaskTracker:
    # 负责执行任务
    def __init__(self):
        self.name = 'cubietruck-plus'
        self.running = []
        pass

    # 扫描数据库看是否有属于自己的任务
    def execute(self, session):
        waiting_task = session.query(TaskInstance) \
            .filter_by(execute_machine=self.name) \
            .filter_by(status='waiting') \
            .all()
        for taskInstance in waiting_task:
            task = session.query(Task).filter_by(id=taskInstance.task_id).first()
            shell = Shell(task.command, task.id, taskInstance.version)
            self.running.append(shell)
            shell.run()
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
            if not shell.is_running():
                print 'running'
                continue
            task_instance = session.query(TaskInstance) \
                .filter_by(task_id=shell.task_id) \
                .filter_by(version=shell.version) \
                .first()
            if shell.success():
                print 'success'
                task_instance.status = 'success'
            else:
                print 'failed'
                task_instance.status = 'failed'
            session.add(task_instance)
        session.commit()

    # 反馈自身负载情况
    def health_feedback(self):
        pass

    def serve(self):
        while True:
            session = DBSession()
            self.execute(session)
            print 'epoch'
            time.sleep(20)
            self.kill(session)
            self.track(session)
            session.close()
            time.sleep(10)


task_tracker = TaskTracker()
task_tracker.serve()
