# coding=utf-8
from . import admin
from flask import render_template, request, redirect
from model import Machine, Action, DBSession
from flask.ext.login import login_required, current_user
from datetime import datetime


@admin.route('/list_machine')
@login_required
def list_machine():
    return render_template('machine/list.html')


@admin.route('/list_machine_status')
@login_required
def list_machine_status():
    return render_template('machine/list_status.html')


@admin.route('/new_machine', methods=['GET', 'POST'])
@login_required
def new_machine():
    if request.method == 'GET':
        return render_template('machine/new.html')
    else:
        print 'hello'
        machine = Machine()
        machine.name = request.form.get('name')
        machine.ip = request.form.get('ip')
        machine.mac = request.form.get('mac')
        session = DBSession()
        session.add(machine)
        session.commit()

        action = Action(user_name=current_user.name, content='新建机器 {}'.format(machine.name), create_time=datetime.now())
        session.add(action)
        session.commit()

        session.close()
        return redirect('/list_machine')


@admin.route('/modify_machine', methods=['POST', 'GET'])
@login_required
def modify_machine():
    if request.method == 'GET':
        extend_id = request.args.get('extend_id')
        return render_template('machine/modify.html')
    else:
        pass
