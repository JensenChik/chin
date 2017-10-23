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

func (log *Log) DumpToMySQL() (bool, error) {
    ok, err := DumpToMySQL(log)
    return ok, err
}

func (log *Log) LoadByWhere(filters ...interface{}) (*Log, error) {
    initLog, err := LoadByWhere(log, filters...)
    if err != nil {
        return nil, err
    } else {
        return initLog.(*Log), nil
    }
}

func (log *Log) LoadByKey(key interface{}) (*Log, error) {
    initLog, err := LoadByKey(log, key)
    if err != nil {
        return nil, err
    } else {
        return initLog.(*Log), nil
    }
}

