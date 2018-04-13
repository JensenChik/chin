package model

import (
    "github.com/jinzhu/gorm"
    "github.com/sdbaiguanghe/glog"
    "../tools/secure"
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
    user.Password = secure.MD5(user.Password)
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

func (user *User) Exists() bool {
    userInDB, err := new(User).LoadByWhere("user_name = ?", user.UserName)
    if err != nil {
        return false
    } else {
        return userInDB.Password == secure.MD5(user.Password)
    }
}

func (user *User) SendMail(message string) bool {
    glog.Error(user.Email)
    glog.Error(message)
    return false
}
