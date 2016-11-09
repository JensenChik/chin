# coding=utf-8
from . import admin
from flask import render_template

@admin.route('/login')
def login():
    return render_template('login.html')

@admin.route('/')
def home():
    return render_template('list_task.html')


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
