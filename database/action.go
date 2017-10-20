package database

import "github.com/jinzhu/gorm"

type Action struct {
    gorm.Model
    UserID  int
    Content string
}

func (action *Action) DumpToMySQL() bool {
    ok := DumpToMySQL(action)
    return ok
}
