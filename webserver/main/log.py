# coding=utf-8
from . import admin
from flask import render_template, request
from model import DBSession, Task, TaskInstance, Action
from flask.ext.login import login_required, current_user
from sqlalchemy import and_, desc
from sqlalchemy.sql.expression import case
from datetime import datetime


@admin.route('/list_execute_log')
@login_required
def list_execute_log():
    session = DBSession()
    tasks_instance = session.query(TaskInstance, Task) \
        .join(Task, and_(TaskInstance.task_id == Task.id, TaskInstance.status.isnot(None))) \
        .order_by(desc(case((
            (and_(TaskInstance.finish_time >= TaskInstance.begin_time, TaskInstance.finish_time >= TaskInstance.pooled_time), TaskInstance.finish_time),
            (and_(TaskInstance.begin_time >= TaskInstance.finish_time, TaskInstance.begin_time >= TaskInstance.pooled_time), TaskInstance.begin_time),
            (and_(TaskInstance.pooled_time >= TaskInstance.finish_time, TaskInstance.pooled_time >= TaskInstance.begin_time), TaskInstance.pooled_time),
        ))))
    session.close()
    return render_template('log/list_execute.html', tasks_instance=tasks_instance)


@admin.route('/list_instance_log')
@login_required
def list_instance_log():
    task_id = int(request.args.get('task_id') or -404)
    session = DBSession()
    tasks_instance = session.query(TaskInstance, Task) \
        .join(Task, and_(TaskInstance.task_id == Task.id,
                         TaskInstance.status.isnot(None),
                         TaskInstance.task_id == task_id)) \
        .order_by(TaskInstance.finish_time.desc(), TaskInstance.begin_time.desc(),
                  TaskInstance.pooled_time.desc()).all()
    session.close()
    return render_template('log/list_instance.html', tasks_instance=tasks_instance)


@admin.route('/list_action_log')
@login_required
def list_action_log():
    session = DBSession()
    actions = session.query(Action).order_by(Action.create_time.desc()).all()
    return render_template('log/list_action.html', actions=actions)


@admin.route('/rerun', methods=['POST'])
@login_required
def rerun():
    task_id = request.form.get('task_id')
    version = request.form.get('version')
    # todo:非法重跑检测
    session = DBSession()
    task_instance = session.query(TaskInstance).filter_by(task_id=task_id, version=version).first()
    task_instance.status = None
    session.commit()

    action = Action(user_name=current_user.name, content='重新执行版本号为 {} 的任务 {}'.format(version, task_id),
                    create_time=datetime.now())
    session.add(action)
    session.commit()

    session.close()
    return render_template('log/list_action.html')


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
