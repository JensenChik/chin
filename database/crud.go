package database

import "github.com/sdbaiguanghe/glog"

func DumpToMySQL(object interface{}) bool {
    db, connectError := ConnectDatabase()
    defer db.Close()
    if connectError != nil {
        return false
    }
    create := db.Create(object)
    if create.Error != nil {
        glog.Errorf("插入 %T 记录失败, ", object)
        glog.Error(create.Error)
        return false
    } else {
        return true
    }
}

