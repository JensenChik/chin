package database

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    UserName    string `gorm:"unique"`
    Password    string
    Email       string
    tmpPassword string
}

func (user *User) BeforeSave(scope *gorm.Scope) error {
    user.tmpPassword = user.Password
    user.Password = toMD5(user.Password)
    return nil
}

func (user *User) AfterSave(scope *gorm.Scope) error {
    user.Password = user.tmpPassword
    return nil
}

func (user *User) DumpToMySQL() (bool, error) {
    ok, err := dumpToMysql(user)
    return ok, err
}

func (user *User) LoadByWhere(filters ...interface{}) (*User, error) {
    initUser, err := loadByWhere(user, filters...)
    if err != nil {
        return nil, err
    } else {
        return initUser.(*User), nil
    }
}

func (user *User) LoadByKey(key interface{}) (*User, error) {
    initUser, err := loadByKey(user, key)
    if err != nil {
        return nil, err
    } else {
        return initUser.(*User), nil
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
