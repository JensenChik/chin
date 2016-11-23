# coding=utf-8
from . import admin
from flask import render_template, request
from flask.ext.login import login_required


@admin.route('/list_machine')
@login_required
def list_machine():
    return render_template('machine/list.html')


@admin.route('/list_machine_status')
@login_required
def list_machine_status():
    return render_template('machine/list_status.html')


@admin.route('/new_machine')
@login_required
def new_machine():
    return render_template('machine/new.html')


@admin.route('/modify_machine', methods=['POST', 'GET'])
@login_required
def modify_machine():
    if request.method == 'GET':
        extend_id = request.args.get('extend_id')
        return render_template('machine/modify.html')
    else:
        pass
