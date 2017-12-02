package database

import "github.com/jinzhu/gorm"

type Operation struct {
    gorm.Model
    UserID  int
    Content string
}

func (op *Operation) DumpToMySQL() (bool, error) {
    ok, err := DumpToMySQL(op)
    return ok, err
}

func (op *Operation) LoadByWhere(filters ...interface{}) (*Operation, error) {
    initOperation, err := LoadByWhere(op, filters...)
    if err != nil {
        return nil, err
    } else {
        return initOperation.(*Operation), nil
    }
}

func (op *Operation) LoadByKey(key interface{}) (*Operation, error) {
    initOperation, err := LoadByKey(op, key)
    if err != nil {
        return nil, err
    } else {
        return initOperation.(*Operation), nil
    }
}
