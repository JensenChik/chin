# coding=utf-8
import time
from datetime import datetime
from model import Task, TaskInstance, DBSession
import ConfigParser


class JobTracker:
    # 负责分配任务，任务调度
    def __init__(self, current_date=None):
        self.current_date = current_date
        cf = ConfigParser.ConfigParser()
        cf.read('chin.ini')
        self.heartbeat_sec = int(cf.get('scheduler', 'heartbeat_sec'))

    # 每天凌晨初始化版本号
    def init_every_day(self, session):
        current_date = datetime.date(datetime.now())
        # 跨天
        if self.current_date != current_date:
            self.current_date = current_date
            valid_tasks = session.query(Task).filter_by(valid=True).all()
            for task in valid_tasks:
                # 当天的版本号
                version = datetime(self.current_date.year, self.current_date.month, self.current_date.day,
                                   task.hour, task.minute, 0).strftime('%Y%m%d%H%M%S')
                # 符合触发条件
                if task.scheduled_type == 'day' or \
                        (task.scheduled_type == 'week' and task.weekday == self.current_date.isoweekday()) or \
                        (task.scheduled_type == 'month' and task.day == self.current_date.day) or \
                        (task.scheduled_type == 'once' and datetime(task.year, task.month,
                                                                    task.day).date() == self.current_date):
                    # 版本号尚未入库
                    if session.query(TaskInstance).filter_by(task_id=task.id, version=version).all() == []:
                        session.add(TaskInstance(task_id=task.id, version=version))
            session.commit()

    # 将可执行的任务置为等待
    def make_waiting(self, session):
        midnight = datetime.date(datetime.now()).strftime("%Y%m%d%H%M%S")
        now = datetime.now().strftime("%Y%m%d%H%M%S")

        # 现在时间大于调度时间的今天任务
        prepared_tasks = session.query(TaskInstance) \
            .filter(TaskInstance.version >= midnight) \
            .filter(TaskInstance.version <= now) \
            .filter(TaskInstance.status == None) \
            .all()

        for task_instance in prepared_tasks:
            # todo:父任务已执行完毕
            task_instance.status = 'waiting'
            task_instance.pooled_time = datetime.now()
            session.add(task_instance)

        session.commit()

    # 为等待中的任务分配机器
    def allocate_machine(self, session):
        waiting_tasks = session.query(TaskInstance).filter_by(status='waiting').all()
        for task_instance in waiting_tasks:
            task_instance.execute_machine = 'cubietruck-plus'
            session.add(task_instance)
        session.commit()

    # 任务失败反馈
    def execute_status_feedback(self, session):
        failed_tasks = session.query(TaskInstance).filter_by(status='failed').filter_by(notify=False).all()
        pass

    def serve(self):
        while True:
            session = DBSession()
            self.init_every_day(session)
            self.make_waiting(session)
            self.allocate_machine(session)
            self.execute_status_feedback(session)
            session.close()
            time.sleep(self.heartbeat_sec)
