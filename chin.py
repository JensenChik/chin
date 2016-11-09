#!/home/suit/dev/anaconda2/bin/python
# coding=utf-8
from core.scheduler import JobTracker
from core.worker import TaskTracker
from model import reset_db
import sys
if __name__ == "__main__":
    action = sys.argv[1]
    if action == "webserver":
        print "star webserver"
    elif action == "worker":
        worker = TaskTracker()
        worker.serve()
    elif action == "scheduler":
        scheduler = JobTracker()
        scheduler.serve()
    elif action == "resetdb":
        reset_db()
    else:
        print "不支持参数%s" % action
