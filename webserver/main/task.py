# coding=utf-8
from . import admin
from flask import render_template, request
from model import DBSession, Task
from flask.ext.login import login_required, current_user
from datetime import datetime
from sqlalchemy.orm.attributes import flag_modified


@admin.route('/list_task')
@login_required
def list_task():
    session = DBSession()
    tasks = session.query(Task).order_by(Task.id.desc()).all()
    session.close()
    return render_template('task/list.html', tasks=tasks)


@admin.route('/new_task', methods=['POST', 'GET'])
@login_required
def new_task():
    if request.method == 'GET':
        extend_id = request.args.get('extend_id')
        return render_template('task/new.html')
    else:
        data = request.form
        father_tasks = data.get('father_task').strip()
        task = Task(
                name=data.get('task_name'),
                user=current_user.name,
                create_time=datetime.now(),
                command=data.get('command'),
                priority=data.get('priority'),
                machine_pool=data.get('machine_pool').split('\n'),
                father_task=[int(i) for i in father_tasks.split('\n')] if father_tasks != '' else [],
                valid=data.get('valid') == 'true',
                rerun=data.get('rerun') == 'true',
                rerun_times=data.get('rerun_times'),
                scheduled_type=data.get('scheduled_type'),
                year=data.get('year'),
                month=data.get('month'),
                weekday=data.get('weekday'),
                day=data.get('day'),
                hour=data.get('hour'),
                minute=data.get('minute')
        )
        session = DBSession()
        session.add(task)

        # 填充父任务的子任务
        # todo:父任务填写错误返回异常
        for father_id in task.father_task:
            father_task = session.query(Task).filter_by(id=father_id).first()
            father_task.child_task.append(task.id)
            flag_modified(father_task, "child_task")

        session.commit()
        session.close()
        return '/list_task' if request.form.get('has_next') == 'false' else 'new_task'


@admin.route('/modify_task', methods=['POST', 'GET'])
@login_required
def modify_task():
    if request.method == 'GET':
        extend_id = request.args.get('extend_id')
        return render_template('task/modify.html')
    else:
        data = request.form
        task_id = int(data.get('task_id'))
        new_father_tasks = data.get('father_task').strip()
        session = DBSession()
        task = session.query(Task).filter_by(id=task_id).first()

        task.name = data.get('task_name')
        task.command = data.get('command')
        task.priority = data.get('priority')
        task.machine_pool = data.get('machine_pool').split('\n')
        task.valid = data.get('valid') == 'true'
        task.rerun = data.get('rerun') == 'true'
        task.rerun_times = data.get('rerun_times')
        task.scheduled_type = data.get('scheduled_type')
        # todo:版本号问题
        task.year = data.get('year')
        task.month = data.get('month')
        task.weekday = data.get('weekday')
        task.day = data.get('day')
        task.hour = data.get('hour')
        task.minute = data.get('minute')

        # 旧的父任务解绑
        for old_father_id in task.father_task:
            old_father_task = session.query(Task).filter_by(id=old_father_id).first()
            old_father_task.child_task.remove(task_id)
            flag_modified(old_father_task, "child_task")

        task.father_task = new_father_tasks.split('\n') if new_father_tasks != '' else []

        # 新父任务绑定
        for father_id in task.father_task:
            father_task = session.query(Task).filter_by(id=father_id).first()
            father_task.child_task.append(task_id)
            flag_modified(father_task, "child_task")

        print 'done'

        session.commit()
        session.close()
        return '/list_task' if request.form.get('has_next') == 'false' else 'modify_task'


@admin.route('/get_task_detail', methods=['POST'])
@login_required
def get_task_detail():
    task_id = int(request.form.get("task_id"))
    session = DBSession()
    task = session.query(Task).filter_by(id=task_id).first()
    session.close()
    return task.to_json()
