package database

import "github.com/jinzhu/gorm"

type Operation struct {
    gorm.Model
    UserID  uint
    Content string
}

func (op *Operation) DumpToMySQL() (bool, error) {
    ok, err := dumpToMysql(op)
    return ok, err
}

func (op *Operation) LoadByWhere(filters ...interface{}) (*Operation, error) {
    initOperation, err := loadByWhere(op, filters...)
    if err != nil {
        return nil, err
    } else {
        return initOperation.(*Operation), nil
    }
}

func (op *Operation) LoadByKey(key interface{}) (*Operation, error) {
    initOperation, err := loadByKey(op, key)
    if err != nil {
        return nil, err
    } else {
        return initOperation.(*Operation), nil
    }
}
