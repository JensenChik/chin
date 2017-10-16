package database

import "github.com/jinzhu/gorm"

type User struct {
    gorm.Model
    UserName string `gorm:"unique"`
    Password string
    Email    string
}

func ExistsUser(userName string, password string) bool {
    return false
}
