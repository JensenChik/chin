# coding=utf-8
from . import admin
from flask import render_template, request
from webserver import logger
import traceback


@admin.app_errorhandler(404)
def not_found(e):
    exstr = traceback.format_exc()
    logger.error(exstr)
    logger.error(request.remote_addr + '请求url' + request.url + '失败, 返回404')
    return render_template('error.html', error_code='404'), 404


@admin.app_errorhandler(500)
def server_error(e):
    exstr = traceback.format_exc()
    logger.error(exstr)
    return render_template('error.html', error_code='500'), 500
