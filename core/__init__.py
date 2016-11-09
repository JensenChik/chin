
# coding=utf-8
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from model import BaseModel, Task, TaskInstance

DATABASE_URI = 'mysql+mysqldb://root:qijinxiu@localhost/chin?charset=utf8'

engine = create_engine(DATABASE_URI, encoding='utf-8')
DBSession = sessionmaker(engine)

# # todo:删掉这行
# BaseModel.metadata.drop_all(engine)
#
# BaseModel.metadata.create_all(engine)
#
# # todo: 删掉后面的
# session = DBSession()
# session.add(Task(name='python 脚本', user='jinxiu.qi', valid=True, group='订单中心', command='python ordercenter.py', scheduled_type='day', hour=0, minute=1))
# session.add(Task(name='shell  脚本', user='jinxiu.qi', valid=True, group='支付中心', command='shell  paycenter.py  ', scheduled_type='day', hour=0, minute=2))
# session.commit()

