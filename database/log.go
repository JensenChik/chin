package database

import "github.com/jinzhu/gorm"

type Log struct {
    gorm.Model
    InstanceID int `gorm:"index"`
    MachineID  int
    Output     string
}
