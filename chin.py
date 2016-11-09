#!/home/suit/dev/anaconda2/bin/python
# coding=utf-8
from core import TaskTracker, JobTracker
import sys
if __name__ == "__main__":
    action = sys.argv[1]
    if action == "webserver":
        print "star webserver"
    elif action == "worker":
        print "start worker"
    elif action == "scheduler":
        print "start scheduler"
    elif action == "initdb":
        print "init db"
    elif action == "resetdb":
        print "resetdb"
    else:
        print "不支持参数%s" % action
