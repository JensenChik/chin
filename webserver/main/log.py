# coding=utf-8
from . import admin
from flask import render_template, request
from model import DBSession, Task, TaskInstance
from flask.ext.login import login_required
from sqlalchemy import and_


@admin.route('/list_execute_log')
@login_required
def list_execute_log():
    session = DBSession()
    tasks_instance = session.query(TaskInstance, Task) \
        .join(Task, and_(TaskInstance.task_id == Task.id, TaskInstance.status.isnot(None))) \
        .order_by(TaskInstance.finish_time.desc(), TaskInstance.begin_time.desc(),
                  TaskInstance.pooled_time.desc()).all()
    session.close()
    return render_template('log/list_execute.html', tasks_instance=tasks_instance)


@admin.route('/get_log_detail', methods=['POST'])
@login_required
def get_log_detail():
    session = DBSession()
    task_id = request.form.get('task_id')
    version = request.form.get('version')
    log_detail = session.query(TaskInstance.log).filter_by(task_id=task_id,
                                                           version=version).first() if version is not None \
        else session.query(TaskInstance.log).filter_by(task_id=task_id).order_by(TaskInstance.begin_time.desc()).first()
    return log_detail


@admin.route('/list_action_log')
@login_required
def list_action_log():
    return render_template('log/list_action.html')
