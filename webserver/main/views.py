# coding=utf-8
from . import admin
from flask import render_template
from model import DBSession, Task

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


@admin.route('/new_task')
def new_task():
    return render_template('new_task.html')


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
