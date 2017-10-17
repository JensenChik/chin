package database

import (
    "github.com/jinzhu/gorm"
    "log"
)

type User struct {
    gorm.Model
    UserName string `gorm:"unique"`
    Password string
    Email    string
}

func (user User) DumpToMySQL() bool {
    db, connectError := ConnectDatabase()
    defer db.Close()
    if connectError != nil {
        return false
    }
    create := db.Create(&user)
    if create.Error != nil {
        log.Print("[Error] 插入user记录失败, ", create.Error)
        return false
    } else {
        return true
    }
}

func ExistsUser(userName string, password string) bool {
    md5Passwd := toMD5(password)
    db, _ := ConnectDatabase()
    defer db.Close()
    user := User{}
    db.Where("user_name =?", userName).First(&user)
    return user.Password == md5Passwd
}
