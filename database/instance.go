package database

import (
    "github.com/jinzhu/gorm"
)

type Instance struct {
    gorm.Model
    JobID     int `gorm:"index"`
    MachineID int
    StdOut    string `gorm:"type:longblob"`
}

func (instance *Instance) BeforeSave(scope *gorm.Scope) error {
    instance.StdOut = zip(instance.StdOut)
    return nil
}

func (instance *Instance) AfterSave(scope *gorm.Scope) error {
    instance.StdOut = unzip(instance.StdOut)
    return nil
}

func (instance *Instance) AfterFind(scope *gorm.Scope) error {
    instance.StdOut = unzip(instance.StdOut)
    return nil
}

func (instance *Instance) DumpToMySQL() (bool, error) {
    ok, err := dumpToMysql(instance)
    return ok, err
}

func (instance *Instance) LoadByWhere(filters ...interface{}) (*Instance, error) {
    initInstance, err := loadByWhere(instance, filters...)
    if err != nil {
        return nil, err
    } else {
        return initInstance.(*Instance), nil
    }
}

func (instance *Instance) LoadByKey(key interface{}) (*Instance, error) {
    initInstance, err := loadByKey(instance, key)
    if err != nil {
        return nil, err
    } else {
        return initInstance.(*Instance), nil
    }
}

