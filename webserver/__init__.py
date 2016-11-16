# -*- coding:utf-8 -*-
__author__ = 'jinxiu.qi'
from flask import Flask
from flask.ext.login import LoginManager
import ConfigParser
import logging

cf = ConfigParser.ConfigParser()
cf.read('chin.ini')

logging.basicConfig(
        level=logging.DEBUG,
        format='[%(levelname)s]\t%(asctime)s\t%(message)s',
        datefmt='%Y-%m-%d %H:%M:%S',
        filename=cf.get('webserver', 'log_path') + '/webserver.log',
        filemode='a'
)
logger = logging.getLogger('webserver')

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
