package scheduler

import (
    "../../database"
    "github.com/sdbaiguanghe/glog"
    "time"
)

func jobTracker() {
    glog.Info("job tracker 开始工作")
    for {
        jobs := []database.Job{}
        database.Fill(jobs).Where("date(created_at) = ?", time.Now().Format("2016-01-02"))
        for _, job := range jobs {
            if job.GetReady() {
                job.CreateInstance()
            } else if job.Finish() {
                job.NotifyIfNeed()
                job.RerunIfNeed()
            }
        }
        time.Sleep(time.Second)
    }
}