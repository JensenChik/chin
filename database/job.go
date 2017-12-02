package database

import "github.com/jinzhu/gorm"

type Job struct {
    gorm.Model
    TaskID int `gorm:"index"`
    Status string
}

func (job *Job) DumpToMySQL() (bool, error) {
    ok, err := DumpToMySQL(job)
    return ok, err
}

func (job *Job) LoadByWhere(filters ...interface{}) (*Job, error) {
    initJob, err := LoadByWhere(job, filters...)
    if err != nil {
        return nil, err
    } else {
        return initJob.(*Job), nil
    }
}

func (job *Job) LoadByKey(key interface{}) (*Job, error) {
    initJob, err := LoadByKey(job, key)
    if err != nil {
        return nil, err
    } else {
        return initJob.(*Job), nil
    }
}
