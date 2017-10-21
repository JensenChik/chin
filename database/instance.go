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
