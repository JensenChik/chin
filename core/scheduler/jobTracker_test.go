package scheduler

import (
    "testing"
    "time"

    . "github.com/franela/goblin"
    "github.com/sdbaiguanghe/glog"
)

func TestJobTracker(t *testing.T) {
    g := Goblin(t)

    g.Describe("job tracker 单元测试", func() {
        s, _ := time.Parse("15:04:05", "00:01:02")
        glog.Error(s)
        t, _ := time.Parse("15:04:05", time.Now().Format("15:04:05"))
        glog.Error(t)
        glog.Error(s.Sub(t) > 0)

    })
}
