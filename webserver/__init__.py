# -*- coding:utf-8 -*-
__author__ = 'jinxiu.qi'
from flask import Flask
from flask.ext.login import LoginManager
import ConfigParser

cf = ConfigParser.ConfigParser()
cf.read('chin.ini')
login_manager = LoginManager()
login_manager.session_protection = 'strong'


def create_app():
    app = Flask(__name__)
    app.config['SECRET_KEY'] = cf.get('webserver', 'secret_key')
    login_manager.init_app(app)

    # 初始化蓝图
    from .main import admin
    app.register_blueprint(admin)

    return app
