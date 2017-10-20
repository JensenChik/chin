package database

import (
    "github.com/jinzhu/gorm"
)

type Log struct {
    gorm.Model
    InstanceID int `gorm:"index"`
    MachineID  int
    Output     string `gorm:"type:longblob"`
}

func (log *Log) BeforeSave(scope *gorm.Scope) error {
    log.Output = zip(log.Output)
    return nil
}

func (log *Log) AfterFind(scope *gorm.Scope) error {
    log.Output = unzip(log.Output)
    return nil
}

func (log *Log) DumpToMySQL() bool {
    ok := DumpToMySQL(log)
    return ok
}

