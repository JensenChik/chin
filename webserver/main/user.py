# coding=utf-8
from . import admin
from flask import render_template, request, redirect
from model import DBSession, User, Action
from flask.ext.login import login_required, current_user
from datetime import datetime
import json


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
        user_name = request.form.get('user_name')
        old_password = request.form.get('old_password')
        new_password = request.form.get('new_password')
        email = request.form.get('email')
        session = DBSession()
        user = session.query(User).filter_by(name=user_name).first()
        if user is not None and user.verify_password(old_password):
            user.password = new_password
            user.email = email
            session.commit()
            action = Action(user_name=current_user.name, content='修改用户 {}'.format(user.name),
                            create_time=datetime.now())
            session.add(action)
            session.commit()
            result = {"status": "success", "data": {"href": "/list_user"}}
        else:
            result = {"status": "failed", "err_info": "不存在用户{}".format(user.name) if user is None else "用户原密码错误"}
        session.close()
        return json.dumps(result)


@admin.route('/get_user_detail', methods=['POST'])
@login_required
def get_user_detail():
    user_name = request.form.get('user_name')
    session = DBSession()
    user = session.query(User).filter_by(name=user_name).first()
    if user is not None:
        return json.dumps({
            "status": "success",
            "data": {
                "email": user.email
            }
        })
    else:
        return json.dumps({
            "status": "failed",
            "data": {},
            "err_info": "不存在用户%s" % user_name
        })
