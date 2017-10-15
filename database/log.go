package database

import "github.com/jinzhu/gorm"

type Log struct {
    gorm.Model
    Instance_id int `gorm:"index"`
    Machine_id  int
    Output      string
}
