# -*- coding:utf-8 -*-
__author__ = 'jinxiu.qi'
from flask import Flask


def create_app():
    app = Flask(__name__)

    # 初始化蓝图
    from .main import admin
    app.register_blueprint(admin)

    return app
