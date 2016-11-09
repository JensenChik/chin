# coding=utf-8
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from model import Task, TaskInstance, BaseModel
import ConfigParser
from datetime import datetime
import json

cf = ConfigParser.ConfigParser()
cf.read('chin.ini')
DATABASE_URI = cf.get('db', 'sql_conn')
engine = create_engine(DATABASE_URI, encoding='utf-8')
DBSession = sessionmaker(engine)


def reset_db():
    BaseModel.metadata.drop_all(engine)
    BaseModel.metadata.create_all(engine)
    session = DBSession()
    session.add(Task(name='python 脚本', user='chin',
                     valid=True, group='python task',
                     create_time=datetime.now(),
                     command='python -c "print 12306"',
                     priority=10, machine_pool=["cubieboard", "arduino"],
                     rerun=True, rerun_times=3,
                     scheduled_type='day', hour=0, minute=1))
    session.add(Task(name='shell 脚本', user='chin',
                     valid=True, group='shell task',
                     create_time=datetime.now(),
                     command='sh -c "echo hello shell"',
                     priority=10, machine_pool=["cubieboard", "arduino"],
                     rerun=True, rerun_times=3,
                     scheduled_type='day', hour=0, minute=2))
    session.commit()
    session.close()
