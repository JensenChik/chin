package database

import "github.com/jinzhu/gorm"

type Action struct {
    gorm.Model
    UserID  int
    Content string
}
