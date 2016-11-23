# coding=utf-8
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from model import Task, TaskInstance, BaseModel, User
import ConfigParser
from datetime import datetime
import json

cf = ConfigParser.ConfigParser()
cf.read('chin.ini')
DATABASE_URI = cf.get('db', 'sql_conn')
engine = create_engine(DATABASE_URI, pool_recycle=3600, encoding='utf-8')
DBSession = sessionmaker(engine)
root_name = cf.get('webserver', 'root_name')
root_password = cf.get('webserver', 'root_password')
root_email = cf.get('webserver', 'root_email')


def mock_db():
    BaseModel.metadata.drop_all(engine)
    BaseModel.metadata.create_all(engine)
    session = DBSession()
    task_param = [(0, 1, [], [2]), (0, 2, [1, 3]), (0, 30, [], [2]), (0, 40, [5, 6]), (4, 12, [], [4]), (5, 36, [], [4]), (6, 18), (7, 46), (8, 36),
            (9, 24), (10, 48), (11, 49), (12, 13), (13, 0), (14, 19), (15, 33), (16, 17),
            (17, 54), (18, 35), (19, 27), (20, 22), (21, 41), (22, 49), (23, 11)]
    for t in task_param:
        try:
            father_task = t[2]
        except:
            father_task = []
        try:
            child_task = t[3]
        except:
            child_task = []
        session.add(Task(name='每日%s:%s调度的python脚本' % (t[0], t[1]), user='chin',
                         valid=True,
                         create_time=datetime.now(),
                         command='python -c "print %s;print %s"' % (t[0], t[1]),
                         priority=6, machine_pool=["cubieboard", "arduino"],
                         father_task=father_task,
                         child_task=child_task,
                         rerun=True, rerun_times=2,
                         scheduled_type='day', hour=t[0], minute=t[1]))
    session.add(Task(name='每周调度的shell脚本', user='chin',
                     valid=True,
                     create_time=datetime.now(),
                     command='sh -c "echo schedule every week"',
                     priority=8, machine_pool=["cubieboard", "arduino"],
                     rerun=True, rerun_times=3,
                     scheduled_type='week', weekday=7, hour=0, minute=3))
    session.add(Task(name='每月调度的shell脚本', user='chin',
                     valid=True,
                     create_time=datetime.now(),
                     command='sh -c "echo schedule every month"',
                     priority=7, machine_pool=["cubieboard", "arduino", "alienware"],
                     rerun=True, rerun_times=3,
                     scheduled_type='month', day=2, hour=0, minute=4))
    session.add(Task(name='调度一次的shell脚本', user='chin',
                     valid=False,
                     create_time=datetime.now(),
                     command='sh -c "echo schedule once"',
                     priority=5, machine_pool=["cubieboard"],
                     rerun=False,
                     scheduled_type='once', year=2016, month=8, day=10, hour=0, minute=4))
    root = User()
    root.name = root_name
    root.password = root_password
    root.email = root_email
    session.add(root)
    session.commit()
    session.close()



def clean_db():
    BaseModel.metadata.drop_all(engine)
    BaseModel.metadata.create_all(engine)
    session = DBSession()
    root = User()
    root.name = root_name
    root.password = root_password
    root.email = root_email
    session.add(root)
    session.commit()
    session.close()
