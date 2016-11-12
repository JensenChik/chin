# coding=utf-8
from . import admin
from flask import render_template, request, redirect, url_for
from model import DBSession, Task
import json


@admin.route('/login')
def login():
    return render_template('login.html')


@admin.route('/')
def home():
    session = DBSession()
    tasks = session.query(Task).order_by(Task.id.desc()).all()
    return render_template('list_task.html', tasks=tasks)


@admin.route('/list_task')
def list_task():
    return render_template('list_task.html')


@admin.route('/new_task', methods=['POST', 'GET'])
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
def list_execute_lg():
    return render_template('list_execute_log.html')


@admin.route('/list_action_log')
def list_action_log():
    return render_template('list_action_log.html')


@admin.route('/list_user')
def list_user():
    return render_template('list_user.html')


@admin.route('/new_user')
def new_user():
    return render_template('new_user.html')
