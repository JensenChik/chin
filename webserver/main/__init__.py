# -*- coding:utf-8 -*-
__author__ = 'jinxiu.qi'
from flask import Blueprint
admin = Blueprint('nazgrim', __name__)

from . import views
from . import errors
