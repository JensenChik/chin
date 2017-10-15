package database

import "github.com/jinzhu/gorm"

type Action struct {
    gorm.Model
    User_id int
    Content string
}
