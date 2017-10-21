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
