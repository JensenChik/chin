package worker

import (
    "github.com/sdbaiguanghe/glog"
    "time"
)

func Serve() {
    glog.Error("worker开始工作")
    go instanceTracker()
    for {
        time.Sleep(time.Second)
    }
}
