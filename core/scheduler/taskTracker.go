package scheduler

import (
    "../../database"
    "time"
    "github.com/sdbaiguanghe/glog"
)

func newJobFor(task database.Task) {
    job := new(database.Job)
    job.TaskID = task.ID
    job.Status = "pooling"
    job.DumpToMySQL()
}

func currentDateCompare(date string) (bool, string) {
    today := time.Now().Format("2006-01-02")
    if today == date {
        return true, ""
    } else {
        return false, today
    }
}

func taskTracker() {
    glog.Info("task tracker 服务启动")
    date := ""
    for {
        equal, currentDate := currentDateCompare(date)
        if !equal{
            date = currentDate
            tasks := []database.Task{}
            database.Fill(&tasks).Where("deleted_at is null ")
            for _, task := range tasks {
                if task.ShouldScheduleToday() && task.NoJobToday() {
                    newJobFor(task)
                }
            }
        }
        time.Sleep(time.Second)
    }

}
