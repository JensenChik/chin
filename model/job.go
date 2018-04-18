package model

import (
    "github.com/jinzhu/gorm"
    "github.com/tidwall/gjson"
)

type Job struct {
    gorm.Model
    TaskID   uint `gorm:"index"`
    RunCount uint
    Status   string
    Notified bool
}

func (job *Job) GetReady() bool {
    if job.Status != "pooling" {
        return false
    }
    task := new(Task)
    task.LoadByKey(job.TaskID)
    return task.ShouldScheduleNow()
}

func (job *Job) CreateInstance() {
    machines := []Machine{}
    Fill(machines).Where("alive = ", true)
    machineID := machines[0].ID // 分配机器 todo: 负载均衡
    instance := Instance{
        JobID:     job.ID,
        MachineID: machineID,
    }
    instance.DumpToMySQL()
    job.Status = "waiting"
    job.DumpToMySQL()
}

func (job *Job) Finish() bool {
    return job.Status == "success" || job.Status == "failed"
}

func (job *Job) NotifyIfNeed() {
    message := string(job.TaskID) + string(job.ID) + string(job.Status)
    task := new(Task)
    task.LoadByKey(job.TaskID)
    notifyList := gjson.Parse(task.NotifyList).Array()
    for _, userID := range notifyList {
        user := new(User)
        user.LoadByKey(userID)
        user.SendMail(message)
    }
}

func (job *Job) RerunIfNeed() {
    // 状态是否为 failed
    if job.Status != "failed" {
        return
    }
    task := new(Task)
    task.LoadByKey(job.TaskID)
    if job.RunCount < task.RetryTimes {
        job.Status = "pooling"
    }
    job.DumpToMySQL()
}

func (job *Job) DumpToMySQL() (bool, error) {
    ok, err := dumpToMysql(job)
    return ok, err
}

func (job *Job) LoadByWhere(filters ...interface{}) (*Job, error) {
    initJob, err := loadByWhere(job, filters...)
    if err != nil {
        return nil, err
    } else {
        return initJob.(*Job), nil
    }
}

func (job *Job) LoadByKey(key interface{}) (*Job, error) {
    initJob, err := loadByKey(job, key)
    if err != nil {
        return nil, err
    } else {
        return initJob.(*Job), nil
    }
}
