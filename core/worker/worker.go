package worker

import (
    "time"

    "github.com/sdbaiguanghe/glog"
)

func Serve() {
    glog.Error("worker开始工作")
    go instanceTracker()
    for {
        time.Sleep(time.Second)
    }
}
