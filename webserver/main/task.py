# coding=utf-8
from . import admin
from flask import render_template, request
from model import DBSession, Task, TaskInstance, Action
from flask.ext.login import login_required, current_user
from datetime import datetime
from sqlalchemy.orm.attributes import flag_modified
import json


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

        # todo: 若运行之间晚于当前时间，则补偿一条版本号

        session.commit()

        action = Action(user_name=current_user.name, content='创建任务 {}'.format(task.id), create_time=datetime.now())
        session.add(action)
        session.commit()

        session.close()
        return '/list_task' if request.form.get('has_next') == 'false' else 'new_task'


@admin.route('/reverse_task_valid', methods=['POST'])
@login_required
def reverse_task_valid():
    task_id = int(request.form.get('task_id'))
    valid = request.form.get('valid') == 'True'
    session = DBSession()
    task = session.query(Task).filter_by(id=task_id).first()
    task.valid = not valid
    session.commit()

    action = Action(user_name=current_user.name, content='{}任务 {}'.format('启用' if task.valid else '停用', task.id),
                    create_time=datetime.now())
    session.add(action)
    session.commit()

    session.close()
    return json.dumps({
        "status": "success",
        "data": {
            "url": "/list_task"
        }
    })


@admin.route('/modify_task', methods=['POST', 'GET'])
@login_required
def modify_task():
    if request.method == 'GET':
        task_id = request.args.get('task_id')
        return render_template('task/modify.html', task_id=task_id)
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

        session.commit()

        action = Action(user_name=current_user.name, content='修改任务 {}'.format(task.id), create_time=datetime.now())
        session.add(action)
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


@admin.route('/run_at_once', methods=['POST'])
def run_at_once():
    task_id = int(request.form.get("task_id"))
    version = datetime.now().strftime('%Y%m%d%H%M%S')
    session = DBSession()
    instance = TaskInstance(task_id=task_id, version=version)
    session.add(instance)
    session.commit()

    action = Action(user_name=current_user.name, content='强制执行任务 {}'.format(task_id), create_time=datetime.now())
    session.add(action)
    session.commit()

    session.close()
    return str(task_id)
