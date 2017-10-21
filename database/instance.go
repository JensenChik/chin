package database

import "github.com/jinzhu/gorm"

type Instance struct {
    gorm.Model
    TaskID int `gorm:"index"`
    Status string
}

func (instance *Instance) DumpToMySQL() (bool, error) {
    ok, err := DumpToMySQL(instance)
    return ok, err
}

func (instance *Instance) LoadByWhere(filters ...interface{}) (*Instance, error) {
    initInstance, err := LoadByWhere(instance, filters...)
    if err != nil {
        return nil, err
    } else {
        return initInstance.(*Instance), nil
    }
}

func (instance *Instance) LoadByKey(key interface{}) (*Instance, error) {
    initInstance, err := LoadByKey(instance, key)
    if err != nil {
        return nil, err
    } else {
        return initInstance.(*Instance), nil
    }
}
