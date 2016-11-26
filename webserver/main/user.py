# coding=utf-8
from . import admin
from flask import render_template, request, redirect
from model import DBSession, User, Action
from flask.ext.login import login_required, current_user
from datetime import datetime


@admin.route('/list_user')
@login_required
def list_user():
    session = DBSession()
    users = session.query(User).order_by(User.id.desc()).all()
    session.close()
    return render_template('user/list.html', users=users)


@admin.route('/new_user', methods=['GET', 'POST'])
@login_required
def new_user():
    if request.method == 'GET':
        return render_template('user/new.html')
    else:
        user = User()
        user.name = request.form.get('user_name')
        user.password = request.form.get('password')
        user.email = request.form.get('email')
        session = DBSession()
        session.add(user)
        session.commit()

        action = Action(user_name=current_user.name, content='新建用户 {}'.format(user.name), create_time=datetime.now())
        session.add(action)
        session.commit()

        session.close()
        return redirect('/list_user')


@admin.route('/modify_user', methods=['GET', 'POST'])
@login_required
def modify_user():
    if request.method == 'GET':
        return render_template('user/modify.html')
    else:
        pass
