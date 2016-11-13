# coding=utf-8
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, Integer, Text, Boolean, String, DateTime, SmallInteger, ForeignKey, Enum, TypeDecorator
import enum, json
from flask.ext.login import UserMixin
from werkzeug.security import generate_password_hash, check_password_hash

BaseModel = declarative_base()


class Json(TypeDecorator):
    impl = Text

    def process_bind_param(self, value, dialect):
        return json.dumps(value)

    def process_result_value(self, value, dialect):
        return json.loads(value)


class Task(BaseModel):
    def fields(self,
               id=Column(Integer, primary_key=True, doc="任务id"),
               name=Column(Text, doc="任务名"),
               user=Column(String(32), doc="任务创建者"),
               create_time=Column(DateTime, doc="任务创建时间"),
               command=Column(Text, doc="任务执行命令"),
               priority=Column(SmallInteger, doc="任务优先级"),
               machine_pool=Column(Json, doc="机器池list"),
               father_task=Column(Json, doc="父任务"),
               child_task=Column(Json, doc="子任务"),
               valid=Column(Boolean, index=True, doc="是否调度"),
               rerun=Column(Boolean, index=True, default=False, doc="当失败时是否自动重新执行"),
               rerun_times=Column(SmallInteger, default=0, doc="重新执行次数"),
               scheduled_type=Column(Enum('once', 'day', 'week', 'month'), index=True, doc="调度频率"),
               year=Column(SmallInteger, doc="调度时间-年"),
               month=Column(SmallInteger, doc="调度时间-月"),
               weekday=Column(SmallInteger, doc="调度时间-周几"),
               day=Column(SmallInteger, doc="调度时间-日"),
               hour=Column(SmallInteger, doc="调度时间-时"),
               minute=Column(SmallInteger, doc="调度时间-分")): pass

    # 基础
    __tablename__ = 'task'
    id = Column(Integer, primary_key=True, doc="任务id")
    name = Column(Text, doc="任务名")
    user = Column(String(32), doc="任务创建者")
    create_time = Column(DateTime, doc="任务创建时间")

    # 执行相关
    command = Column(Text, doc="任务执行命令")
    priority = Column(SmallInteger, default=0, doc="任务优先级")
    machine_pool = Column(Json, default=[], doc="机器池list")
    father_task = Column(Json, default=[], doc="父任务")
    child_task = Column(Json, default=[], doc="子任务")

    # 调度相关
    valid = Column(Boolean, index=True, default=False, doc="是否调度")
    rerun = Column(Boolean, index=True, default=False, doc="当失败时是否自动重新执行")
    rerun_times = Column(SmallInteger, default=0, doc="重新执行次数")
    scheduled_type = Column(Enum('once', 'day', 'week', 'month', 'year'), index=True, doc="调度频率")
    year = Column(SmallInteger, doc="调度时间-年")
    month = Column(SmallInteger, doc="调度时间-月")
    weekday = Column(SmallInteger, doc="调度时间-周几")
    day = Column(SmallInteger, doc="调度时间-日")
    hour = Column(SmallInteger, doc="调度时间-时")
    minute = Column(SmallInteger, default=0, doc="调度时间-分")

    def __repr__(self):
        return '<Task %s>' % self.id


class TaskInstance(BaseModel):
    def fields(
            self,
            id=Column(Integer, primary_key=True, doc="日志id"),
            task_id=Column(Integer, ForeignKey('task.id'), doc='任务id'),
            version=Column(String(14), doc='版本号'),
            execute_machine=Column(String(32), doc='执行机器'),
            pooled_time=Column(DateTime, doc='入池时间'),
            begin_time=Column(DateTime, doc='开始执行时间'),
            finish_time=Column(DateTime, doc='执行结束时间'),
            run_count=Column(SmallInteger, default=0, doc="执行次数"),
            status=Column(Enum('waiting', 'abandon', 'running', 'finish', 'failed', 'killing', 'repairing'), index=True,
                          doc='状态'),
            notify=Column(Boolean, index=True, default=False, doc="是否已报警")
    ): pass

    __tablename__ = 'task_instance'

    id = Column(Integer, primary_key=True, doc="日志id")
    task_id = Column(Integer, ForeignKey('task.id'), doc='任务id')
    version = Column(String(14), doc='版本号')
    execute_machine = Column(String(32), doc='执行机器')
    pooled_time = Column(DateTime, doc='入池时间')
    begin_time = Column(DateTime, doc='开始执行时间')
    finish_time = Column(DateTime, doc='执行结束时间')
    run_count = Column(SmallInteger, default=0, doc="执行次数")
    status = Column(Enum('waiting', 'abandon', 'running', 'finish', 'failed', 'killing', 'repairing'), index=True,
                    doc='状态')
    notify = Column(Boolean, index=True, default=False, doc="是否已报警")

    def __repr__(self):
        return '<TaskQueue %s>' % self.id


class User(UserMixin, BaseModel):
    __tablename__ = 'user'
    id = Column(Integer, primary_key=True)
    name = Column(String(64), unique=True)
    password_hash = Column(String(128))
    email = Column(String(64))

    @property
    def password(self):
        raise AttributeError('password 只有写权限')

    @password.setter
    def password(self, password):
        self.password_hash = generate_password_hash(password)

    def verify_password(self, password):
        return check_password_hash(self.password_hash, password)

    def __repr__(self):
        return 'user', self.id, self.name



