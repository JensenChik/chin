package database

import (
    "github.com/jinzhu/gorm"
    "github.com/sdbaiguanghe/glog"
    "log"
    "time"
)

type Log struct {
    gorm.Model
    InstanceID int `gorm:"index"`
    MachineID  int
    Output     string
}

func (log *Log) DumpToMySQL() bool {
    db, connectError := ConnectDatabase()
    defer db.Close()
    if connectError != nil {
        return false
    }
    log.Output = zip(log.Output)
    create := db.Create(&log)
    if create.Error != nil {
        glog.Error("插入 log 记录失败, ", create.Error)
        return false
    } else {
        return true
    }
}




