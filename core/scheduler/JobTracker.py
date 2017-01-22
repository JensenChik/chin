# coding=utf-8
import time
from datetime import datetime
from model import Task, TaskInstance, DBSession
import ConfigParser
import logging
import traceback
import os
from core.scheduler.Email import Email


class JobTracker:
    # 负责分配任务，任务调度
    def __init__(self, current_date=None):
        self.current_date = current_date
        cf = ConfigParser.ConfigParser()
        cf.read('chin.ini')
        self.heartbeat_sec = int(cf.get('scheduler', 'heartbeat_sec'))

        # 配置logger
        self.log_path = cf.get('scheduler', 'log_path')
        self.logger = logging.getLogger('scheduler')
        handler = logging.FileHandler(os.path.join(self.log_path, 'scheduler.log'))
        handler.setLevel(logging.INFO)
        handler.setFormatter(logging.Formatter('[%(levelname)s]\t%(asctime)s\t%(message)s'))
        self.logger.addHandler(handler)

    # 每天凌晨初始化版本号
    def init_every_day(self, session):
        current_date = datetime.date(datetime.now())
        # 跨天
        if self.current_date != current_date:
            self.current_date = current_date

            # 将昨天没有执行的任务置为失败
            undo_tasks = session.query(TaskInstance) \
                .filter(TaskInstance.version < str(self.current_date.strftime('%Y%m%d'))) \
                .filter(TaskInstance.status == None) \
                .all()
            for task_instance in undo_tasks:
                task_instance.log = "任务当天没有执行，被调度器杀死"
                task_instance.status = "failed"

            # 生成当天的版本号
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
            # 如果今天的父任务们都执行完毕
            father_all_done = True
            father_id = session.query(Task.father_task).filter_by(id=task_instance.task_id).first()[0]
            if father_id != []:
                father_status = session.query(TaskInstance.status) \
                    .filter(TaskInstance.task_id.in_(father_id)) \
                    .filter(TaskInstance.version >= midnight).all()
                for status in father_status:
                    if status[0] != 'success':
                        father_all_done = False
            if father_all_done:
                task_instance.status = 'waiting'
                task_instance.pooled_time = datetime.now()

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
        for task_instance in failed_tasks:
            subject = '[任务失败] 任务 {}-{} 执行失败'.format(task_instance.task_id, task_instance.version)
            msg = '''执行机器:{}\n入池时间:{}\n开始时间:{}\n结束时间:{}\n执行次数:{}\n日志详情:\n{}'''.format(
                task_instance.execute_machine,
                task_instance.pooled_time,
                task_instance.begin_time,
                task_instance.finish_time,
                task_instance.run_count,
                task_instance.log
            )
            email = Email()
            email.send(subject, msg)
            task_instance.notify = True
        session.commit()

    def serve(self):
        while True:
            try:
                session = DBSession()
                self.init_every_day(session)
                self.make_waiting(session)
                self.allocate_machine(session)
                self.execute_status_feedback(session)
                session.close()
            except Exception, e:
                self.logger.error(e)
                self.logger.error(traceback.format_exc())
            time.sleep(self.heartbeat_sec)
