package database

import "github.com/jinzhu/gorm"

type Task struct {
    gorm.Model
    TaskName       string
    Command        string
    FatherTask     string `gorm:"type:text"`
    Valid          bool
    MachinePool    string
    OwnerID        int
    ScheduleType   string
    ScheduleFormat string
}

func (task *Task) DumpToMySQL() (bool, error) {
    ok, err := DumpToMySQL(task)
    return ok, err
}

func (task *Task) LoadByWhere(filters ...interface{}) (*Task, error) {
    initTask, err := LoadByWhere(task, filters...)
    if err != nil {
        return nil, err
    } else {
        return initTask.(*Task), nil
    }
}

func (task *Task) LoadByKey(key interface{}) (*Task, error) {
    initTask, err := LoadByKey(task, key)
    if err != nil {
        return nil, err
    } else {
        return initTask.(*Task), nil
    }
}
