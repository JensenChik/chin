package scheduler

import (
    "time"

    "github.com/sdbaiguanghe/glog"
)

func Serve() {
    glog.Info("调度器开始工作")
    go taskTracker()
    go jobTracker()
    for {
        time.Sleep(time.Second)
    }
}
