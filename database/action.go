package database

import "github.com/jinzhu/gorm"

type Action struct {
    gorm.Model
    UserID  int
    Content string
}

func (action *Action) DumpToMySQL() (bool, error) {
    ok, err := DumpToMySQL(action)
    return ok, err
}

func (action *Action) LoadByWhere(filters ...interface{}) (*Action, error) {
    initAction, err := LoadByWhere(action, filters...)
    if err != nil {
        return nil, err
    } else {
        return initAction.(*Action), nil
    }
}

func (action *Action) LoadByKey(key interface{}) (*Action, error) {
    initAction, err := LoadByKey(action, key)
    if err != nil {
        return nil, err
    } else {
        return initAction.(*Action), nil
    }
}
