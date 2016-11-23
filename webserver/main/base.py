# coding=utf-8
from . import admin
from flask import render_template, request, redirect, send_file
from model import DBSession, User
from flask.ext.login import login_user, login_required, logout_user
from webserver import login_manager


@login_manager.user_loader
def load_user(user_id):
    session = DBSession()
    uid = session.query(User).get(int(user_id))
    session.close()
    return uid


@login_manager.unauthorized_handler
def unauthorized():
    return redirect('/login')


@admin.route('/login', methods=['GET', 'POST'])
def login():
    if request.method == 'GET':
        return render_template('base/login.html')
    else:
        user_name = request.form.get('user_name')
        password = request.form.get('password')
        session = DBSession()
        user = session.query(User).filter_by(name=user_name).first()
        session.close()
        if user is not None and user.verify_password(password):
            login_user(user)
        return redirect('/')


@admin.route('/logout')
@login_required
def logout():
    logout_user()
    return redirect('/login')


@admin.route('/')
@login_required
def home():
    return redirect('/list_task')


@admin.route('/favicon.ico')
def ico():
    return send_file('static/chin.ico')
