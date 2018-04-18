package scheduler

import (
    "time"

    "../../model"
    "../../tools/datetime"
    "github.com/sdbaiguanghe/glog"
)

func jobTracker() {
    glog.Info("job tracker 开始工作")
    for {
        jobs := []model.Job{}
        model.Fill(jobs).Where("date(created_at) = ?", datetime.Today())
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
