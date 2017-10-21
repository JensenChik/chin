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
