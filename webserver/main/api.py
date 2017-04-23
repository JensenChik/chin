from . import admin
from flask import request
from flask.ext.login import login_required, current_user
from model import DBSession, Task, TaskInstance, Action
import json


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
