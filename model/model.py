# coding=utf-8
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, Integer, Text, Boolean, String, DateTime, SmallInteger, ForeignKey, Enum, TypeDecorator
import enum, json, zlib
from flask.ext.login import UserMixin
from werkzeug.security import generate_password_hash, check_password_hash
from sqlalchemy.types import LargeBinary
from sqlalchemy.dialects.mysql import LONGBLOB

BaseModel = declarative_base()


class Json(TypeDecorator):
    impl = Text

    def process_bind_param(self, value, dialect):
        return json.dumps(value)

    def process_result_value(self, value, dialect):
        return json.loads(value)


class BinaryString(TypeDecorator):
    # impl = LargeBinary
    impl = LONGBLOB

    def process_bind_param(self, value, dialect):
        return zlib.compress(value)

    def process_result_value(self, value, dialect):
        return zlib.decompress(value)


class Task(BaseModel):
    __tablename__ = 'task'

    # 基础
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
    scheduled_type = Column(Enum('once', 'day', 'week', 'month'), index=True, doc="调度频率")
    year = Column(SmallInteger, doc="调度时间-年")
    month = Column(SmallInteger, doc="调度时间-月")
    weekday = Column(SmallInteger, doc="调度时间-周几")
    day = Column(SmallInteger, doc="调度时间-日")
    hour = Column(SmallInteger, doc="调度时间-时")
    minute = Column(SmallInteger, default=0, doc="调度时间-分")

    def __repr__(self):
        return '<Task %s>' % self.id

    def to_dict(self):
        return {
            'id': self.id,
            'name': self.name,
            'user': self.user,
            'create_time': self.create_time and str(self.create_time),
            'command': self.command,
            'priority': self.priority,
            'machine_pool': self.machine_pool,
            'father_task': self.father_task,
            'child_task': self.child_task,
            'valid': self.valid,
            'rerun': self.rerun,
            'rerun_times': self.rerun_times,
            'scheduled_type': self.scheduled_type,
            'year': self.year,
            'month': self.month,
            'weekday': self.weekday,
            'day': self.day,
            'hour': self.hour,
            'minute': self.minute
        }


class TaskInstance(BaseModel):
    __tablename__ = 'task_instance'

    id = Column(Integer, primary_key=True, doc="日志id")
    task_id = Column(Integer, ForeignKey('task.id'), doc='任务id')
    version = Column(String(14), doc='版本号')
    execute_machine = Column(String(32), doc='执行机器')
    pooled_time = Column(DateTime, doc='入池时间')
    begin_time = Column(DateTime, doc='开始执行时间')
    finish_time = Column(DateTime, doc='执行结束时间')
    run_count = Column(SmallInteger, default=0, doc="执行次数")
    status = Column(Enum('waiting', 'running', 'success', 'failed', 'killing', 'repairing'), index=True,
                    doc='状态')
    log = Column(BinaryString, default='', doc="日志")
    notify = Column(Boolean, index=True, default=False, doc="是否已报警")

    def __repr__(self):
        return '<TaskQueue %s>' % self.id

    def to_dict(self, expected=None, unexpected=None):
        expected = expected or ["id", "task_id", "version", "execute_machine", "pooled_time",
                                "begin_time", "finish_time", "run_count", "status", "log", "notify"]
        unexpected = unexpected or []
        if isinstance(expected, str): expected = [expected]
        if isinstance(unexpected, str): unexpected = [unexpected]
        expected = [col for col in expected if col not in unexpected]
        return dict(
            (col_name, getattr(self, col_name)) if col_name not in ['pooled_time', 'begin_time', 'finish_time']
            else (col_name, str(getattr(self, col_name))) for col_name in expected
        )


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


class Action(BaseModel):
    __tablename__ = 'action'

    id = Column(Integer, primary_key=True, doc="操作id")
    user_name = Column(String(64), doc='用户')
    content = Column(String(128), doc='操作详情')
    create_time = Column(DateTime, doc='操作时间')

    def __repr__(self):
        return '<action %s>' % self.id


class Machine(BaseModel):
    __tablename__ = 'machine'

    id = Column(Integer, primary_key=True, doc="机器id")
    name = Column(String(64), unique=True, index=True, doc="机器名")
    ip = Column(String(15), doc="机器ip")
    mac = Column(String(20), doc="机器mac地址")
    cpu_load = Column(Integer, doc="当前cpu负载")
    men_load = Column(Integer, doc="当前内存负载")
    update_time = Column(DateTime, doc="更新时间")

    def __repr__(self):
        return '<machine %s>' % self.id
