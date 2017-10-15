package database

import "github.com/jinzhu/gorm"

type Instance struct {
    gorm.Model
    Task_id int `gorm:"index"`
    Status  string
}
