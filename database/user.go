package database

import "github.com/jinzhu/gorm"

type User struct {
    gorm.Model
    User_name string `gorm:"unique"`
    Password  string
    Email     string
}
