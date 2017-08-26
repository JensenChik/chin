# -*- coding:utf-8 -*-
__author__ = 'jinxiu.qi'
from flask import Flask
from flask.ext.login import LoginManager
import ConfigParser
import logging
import os

cf = ConfigParser.ConfigParser()
cf.read('chin.ini')

logger = logging.getLogger('webserver')
handler = logging.FileHandler(os.path.join(cf.get('webserver', 'log_path'), 'webserver.log'))
handler.setLevel(logging.DEBUG)
handler.setFormatter(logging.Formatter('[%(levelname)s]\t%(asctime)s\t%(message)s'))
logger.addHandler(handler)

login_manager = LoginManager()
login_manager.session_protection = 'strong'

def zfill(s, length=2):
    return str(s).zfill(length)

def create_app():
    app = Flask(__name__)
    app.config['SECRET_KEY'] = cf.get('webserver', 'secret_key')
    app.add_template_filter(zfill, name='zfill')
    login_manager.init_app(app)

    # 初始化蓝图
    from .main import admin
    app.register_blueprint(admin)

    return app
