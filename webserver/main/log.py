# coding=utf-8
from . import admin
from flask import render_template, request
from model import DBSession, Task, TaskInstance, Action
from flask.ext.login import login_required, current_user
from sqlalchemy import and_, desc
from sqlalchemy.sql.functions import coalesce
from sqlalchemy.sql import func
from datetime import datetime
import json


@admin.route('/list_execute_log')
@login_required
def list_execute_log():
    return render_template('log/list_execute.html')


@admin.route('/list_instance_log')
@login_required
def list_instance_log():
    task_id = int(request.args.get('task_id') or -404)
    session = DBSession()
    tasks_instance = session.query(TaskInstance, Task) \
        .join(Task, and_(TaskInstance.task_id == Task.id,
                         TaskInstance.status.isnot(None),
                         TaskInstance.task_id == task_id)) \
        .order_by(TaskInstance.finish_time.desc(), TaskInstance.begin_time.desc(),
                  TaskInstance.pooled_time.desc()).all()
    session.close()
    return render_template('log/list_instance.html', tasks_instance=tasks_instance)


@admin.route('/list_action_log')
@login_required
def list_action_log():
    session = DBSession()
    actions = session.query(Action).order_by(Action.create_time.desc()).all()
    return render_template('log/list_action.html', actions=actions)


@admin.route('/rerun', methods=['POST'])
@login_required
def rerun():
    task_id = request.form.get('task_id')
    version = request.form.get('version')
    # todo:非法重跑检测
    session = DBSession()
    task_instance = session.query(TaskInstance).filter_by(task_id=task_id, version=version).first()
    task_instance.pooled_time = datetime.now()
    task_instance.status = 'waiting'
    session.commit()

    action = Action(user_name=current_user.name, content='重新执行版本号为 {} 的任务 {}'.format(version, task_id),
                    create_time=datetime.now())
    session.add(action)
    session.commit()

    session.close()
    return render_template('log/list_action.html')


@admin.route('/get_log_by_page', methods=['GET'])
@login_required
def get_log_by_page():
    from jinja2 import Template
    status_template = """
        {% if status == 'waiting' %}
            <a class="btn btn-default btn-xs">等待中</a>
        {% elif status == 'running' %}
            <a class="btn btn-warning btn-xs">运行中</a>
        {% elif status == 'success' %}
            <a class="btn btn-success btn-xs">成功</a>
            <a data-toggle="modal" data-target="#log_detail" class="btn btn-default btn-xs"
               onclick="fill_log_to_div('{{ task_id }}', '{{ version }}', 'log_content')">查看日志</a>
        {% elif status == 'killing' %}
            <a class="btn btn-danger btn-xs">killing</a>
        {% elif status == 'failed' %}
            <a class="btn btn-danger btn-xs">失败</a>
            <a data-toggle="modal" data-target="#log_detail" class="btn btn-default btn-xs"
               onclick="fill_log_to_div('{{ task_id }}', '{{ version }}', 'log_content')">查看日志</a>
        {% elif status == 'repairing' %}
            <a class="btn btn-danger btn-xs">修复中</a>
        {% endif %}
    """

    action_template = """
        <div class="input-group-btn dropdown">
            <button type="button" class="btn btn-default btn-xs dropdown-toggle" data-toggle="dropdown">
                操作<span class="caret"></span></button>
            <ul class="dropdown-menu dropdown-menu-right" role="menu">
                {% if status == 'failed' or status == 'success' %}
                    <li><a onclick="rerun('{{ task_id }}', '{{ version }}')">重跑</a></li>
                {% endif %}
                {% if status == 'running' %}
                    <li><a onclick="kill_task('{{ task_id }}', '{{ version }}')">中止执行</a></li>
                {% elif status == 'failed' or status == 'repairing' %}
                    <li><a href="#">置为成功</a></li>
                    <li><a href="#">置为修复中</a></li>
                {% elif status == 'success' %}
                    <li><a href="#">置为失败</a></li>
                    <li><a href="#">置为修复中</a></li>
                {% endif %}
            </ul>
        </div>
    """

    limit = int(request.args.get('limit'))
    offset = int(request.args.get('offset'))
    session = DBSession()

    tasks_instance = session.query(TaskInstance) \
        .filter(TaskInstance.status != None) \
        .order_by(desc(func.greatest(
        coalesce(TaskInstance.pooled_time, func.date('1900-01-01')),
        coalesce(TaskInstance.begin_time, func.date('1900-01-01')),
        coalesce(TaskInstance.finish_time, func.date('1900-01-01'))))
    )
    count = tasks_instance.count()
    tasks_instance = tasks_instance.offset(offset).limit(limit).all()

    table = []

    for instance in tasks_instance:
        meta = session.query(Task).filter_by(id=instance.task_id).first()
        row = {
            'task_id': instance.task_id,
            'version': instance.version,
            'name': meta.name,
            'execute_machine': '<a class="btn btn-default btn-xs">%s</a>' % instance.execute_machine,
            'pooled_time': str(instance.pooled_time or ''),
            'begin_time': str(instance.begin_time or ''),
            'finish_time': str(instance.finish_time or ''),
            'run_count': instance.run_count,
            'status': Template(status_template).render(
                status=instance.status,
                task_id=instance.task_id,
                version=instance.version
            ),
            'action': Template(action_template).render(
                status=instance.status,
                task_id=instance.task_id,
                version=instance.version
            )
        }
        table.append(row)
    session.close()

    return json.dumps({'total': count, 'rows': table})


@admin.route('/get_log_detail', methods=['POST'])
@login_required
def get_log_detail():
    session = DBSession()
    task_id = request.form.get('task_id')
    version = request.form.get('version')
    log_detail = session.query(TaskInstance.log).filter_by(task_id=task_id,
                                                           version=version).first() if version is not None \
        else session.query(TaskInstance.log).filter_by(task_id=task_id).order_by(TaskInstance.begin_time.desc()).first()
    return log_detail
