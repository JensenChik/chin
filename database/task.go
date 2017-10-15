package database

import "github.com/jinzhu/gorm"

type Task struct {
    gorm.Model
    Task_name       string
    Command         string
    Father_task     string `gorm:"type:text"`
    Valid           bool
    Machine_pool    string
    Owner_id        int
    Schedule_type   string
    Schedule_format string
}
