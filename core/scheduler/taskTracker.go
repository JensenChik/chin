package scheduler

import (
    "../../model"
    "../../tools/datetime"
    "time"
    "github.com/sdbaiguanghe/glog"
)

func taskTracker() {
    glog.Info("task tracker 服务启动")
    date := ""
    for {
        if dt := datetime.Today(); date != dt {
            date = dt
            tasks := []model.Task{}
            model.Fill(&tasks).Where("deleted_at is null ")
            for _, task := range tasks {
                if task.ShouldScheduleToday() && task.NoJobToday() {
                    task.CreateJob()
                }
            }
        }
        time.Sleep(time.Second)
    }

}
