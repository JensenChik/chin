package database

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    UserName string `gorm:"unique"`
    Password string
    Email    string
}

func (user *User) BeforeSave(scope *gorm.Scope) error {
    user.Password = toMD5(user.Password)
    return nil
}

func (user *User) DumpToMySQL() bool {
    ok := DumpToMySQL(user)
    return ok
}

func ExistsUser(userName string, password string) bool {
    md5Passwd := toMD5(password)
    db, _ := ConnectDatabase()
    defer db.Close()
    user := User{}
    db.Where("user_name =?", userName).First(&user)
    return user.Password == md5Passwd
}
