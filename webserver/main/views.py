# coding=utf-8
from . import admin
from flask import render_template, request, redirect, url_for
from model import DBSession, Task, TaskInstance, User
from flask.ext.login import login_user, login_required
from webserver import login_manager
import json


@login_manager.user_loader
def load_user(user_id):
    session = DBSession()
    return session.query(User).get(int(user_id))


@login_manager.unauthorized_handler
def unauthorized():
    return redirect('/login')


@admin.route('/login', methods=['GET', 'POST'])
def login():
    if request.method == 'GET':
        return render_template('login.html')
    else:
        user_name = request.form.get('user_name')
        password = request.form.get('password')
        session = DBSession()
        user = session.query(User).filter_by(name=user_name).first()
        session.close()
        if user is not None and user.verify_password(password):
            login_user(user)
        return redirect('/')


@admin.route('/')
@login_required
def home():
    session = DBSession()
    tasks = session.query(Task).order_by(Task.id.desc()).all()
    session.close()
    return render_template('list_task.html', tasks=tasks)


@admin.route('/list_task')
@login_required
def list_task():
    return render_template('list_task.html')


@admin.route('/new_task', methods=['POST', 'GET'])
@login_required
def new_task():
    if request.method == 'GET':
        extend_id = request.args.get('extend_id')
        return render_template('new_task.html')
    else:
        task_name = request.form.get('task_name')
        command = request.form.get('command')
        valid = request.form.get('valid')
        priority = request.form.get('priority')
        rerun = request.form.get('rerun')
        rerun_time = request.form.get('rerun_times')
        father_task = request.form.get('father_task')
        machine_pool = request.form.get('machine_pool')
        return '/' if request.form.get('has_next') == 'false' else 'new_task'


@admin.route('/list_execute_log')
@login_required
def list_execute_lg():
    session = DBSession()
    tasks_instance = session.query(TaskInstance, Task) \
        .join(Task, TaskInstance.task_id == Task.id) \
        .order_by(TaskInstance.id.desc()).all()
    session.close()
    return render_template('list_execute_log.html', tasks_instance=tasks_instance)


@admin.route('/list_action_log')
@login_required
def list_action_log():
    return render_template('list_action_log.html')


@admin.route('/list_user')
@login_required
def list_user():
    return render_template('list_user.html')


@admin.route('/new_user')
@login_required
def new_user():
    return render_template('new_user.html')


@admin.route('/list_machine')
@login_required
def list_machine():
    return render_template('list_machine.html')


@admin.route('/list_machine_status')
@login_required
def list_machine_status():
    return render_template('list_machine_status.html')


@admin.route('/new_machine')
@login_required
def new_machine():
    return render_template('new_machine.html')
