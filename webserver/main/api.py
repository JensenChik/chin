from . import admin
from flask import request
from flask.ext.login import login_required, current_user
from model import DBSession, Task, TaskInstance, Action
import json
from sqlalchemy import and_, desc
from sqlalchemy.sql.functions import coalesce
from sqlalchemy.sql import func


@admin.route('/api/list_task')
@login_required
def api_list_task():
    session = DBSession()
    tasks = session.query(Task).order_by(Task.id.desc()).all()
    session.close()
    return json.dumps({
        "status": "success",
        "data": [task.to_dict() for task in tasks]
    }, ensure_ascii=False)


@admin.route('/api/list_instance_of/<task_id>')
@login_required
def api_list_instance_by(task_id):
    session = DBSession()
    instance = session.query(TaskInstance) \
        .filter(and_(TaskInstance.status != None, TaskInstance.task_id == task_id)) \
        .order_by(desc(func.greatest(
        coalesce(TaskInstance.pooled_time, func.date('1900-01-01')),
        coalesce(TaskInstance.begin_time, func.date('1900-01-01')),
        coalesce(TaskInstance.finish_time, func.date('1900-01-01'))))
    ).all()
    session.close()
    return json.dumps({
        "status": "success",
        "data": [i.to_dict().__delitem__('log') for i in instance]
    }, ensure_ascii=False)


@admin.route('/api/task_detail')
@login_required
def api_task_detail():
    task_id = int(request.args.get('task_id'))
    session = DBSession()
    task = session.query(Task).filter_by(id=task_id).first()
    return json.dumps({
        "status": "success",
        "data": task.to_dict()
    }, ensure_ascii=False)
