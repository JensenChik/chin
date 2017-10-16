package database

import "github.com/jinzhu/gorm"

type Machine struct {
    gorm.Model
    MachineName string
    IP          string `gorm:"size:15"`
    MAC         string `gorm:"size:17"`
    CPULoad     int
    MemoryLoad  int
    Alive       bool
}
