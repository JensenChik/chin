#!/home/suit/dev/anaconda2/bin/python
# -*- coding:utf-8 -*-
__author__ = 'jinxiu.qi'
import sys

reload(sys)
sys.setdefaultencoding('utf-8')

from core.scheduler import JobTracker
from core.worker import TaskTracker
from model import reset_db
from flask.ext.script import Manager, Shell, Server
from webserver import create_app
import ConfigParser

cf = ConfigParser.ConfigParser()
cf.read('chin.ini')
host = cf.get('webserver', 'host')
port = cf.get('webserver', 'port')
app = create_app()
manager = Manager(app)
server = Server(host=host, port=port)
manager.add_command("runserver", server)

if __name__ == "__main__":
    action = sys.argv[1]
    if action == "runserver":
        manager.run()

    elif action == "worker":
        worker = TaskTracker()
        worker.serve()

    elif action == "scheduler":
        scheduler = JobTracker()
        scheduler.serve()

    elif action == "resetdb":
        reset_db()

    else:
        print "不支持参数%s" % action, '只支持参数 runserver worker scheduler resetdb'
