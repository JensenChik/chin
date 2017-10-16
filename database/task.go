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
