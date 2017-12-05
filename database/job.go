package database

import "github.com/jinzhu/gorm"

type Job struct {
    gorm.Model
    TaskID    uint `gorm:"index"`
    MachineID uint `gorm:"index"`
    Status    string
    Notified  bool
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
