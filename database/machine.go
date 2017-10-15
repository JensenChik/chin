package database

import "github.com/jinzhu/gorm"

type Machine struct {
    gorm.Model
    Machine_name string
    IP           string `gorm:"size:15"`
    MAC          string `gorm:"size:17"`
    CPU_load     int
    Memory_load  int
    alive        bool
}
