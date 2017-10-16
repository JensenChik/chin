package database

import "github.com/jinzhu/gorm"

type Instance struct {
    gorm.Model
    TaskID int `gorm:"index"`
    Status string
}
