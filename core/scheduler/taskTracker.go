package scheduler

import (
    "../../database"
    "time"
    "github.com/sdbaiguanghe/glog"
)

func currentDate() (string) {
    return time.Now().Format("2006-01-02")
}

func taskTracker() {
    glog.Info("task tracker 服务启动")
    date := ""
    for {
        if dt := currentDate(); date != dt {
            date = dt
            tasks := []database.Task{}
            database.Fill(&tasks).Where("deleted_at is null ")
            for _, task := range tasks {
                if task.ShouldScheduleToday() && task.NoJobToday() {
                    task.CreateJob()
                }
            }
        }
        time.Sleep(time.Second)
    }

}
