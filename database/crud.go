package database

import (
    "github.com/sdbaiguanghe/glog"
    "errors"
)

func DumpToMySQL(object interface{}) (bool, error) {
    db, connectError := ConnectDatabase()
    defer db.Close()
    if connectError != nil {
        glog.Error(connectError)
        return false, errors.New("无法连接mysql")
    }
    create := db.Save(object)
    if create.Error != nil {
        glog.Errorf("插入 %T 记录失败: %s", object, create.Error)
        return false, errors.New("记录保存失败")
    } else {
        return true, nil
    }
}

func LoadByWhere(object interface{}, filters ...interface{}) (interface{}, error) {
    db, connectError := ConnectDatabase()
    defer db.Close()
    if connectError != nil {
        glog.Error(connectError)
        return false, errors.New("无法连接mysql")
    }
    var recordCount = 0
    checkCount := db.Model(object).Where(filters[0], filters[1:]...).Count(&recordCount)
    if checkCount.Error != nil {
        glog.Error("查询逻辑错误", checkCount.Error)
        return nil, errors.New("查询逻辑错误")
    }
    if recordCount > 1 {
        glog.Errorf("%T中满足条件的记录数为 %d 条而不是 1 条", object, recordCount)
        return nil, errors.New("存在多条满足条件的记录，无法实例化")
    }
    if recordCount == 0 {
        glog.Errorf("%T不存在满足条件的记录，无法实例化", object)
        return nil, errors.New("不存在满足条件的记录，无法实例化")
    }
    load := db.Where(filters[0], filters[1:]...).Find(object).First(object)
    if load.Error != nil {
        glog.Errorf("查询 %T 记录失败, %s", object, load.Error)
        return nil, errors.New("查询失败")
    } else {
        return object, nil
    }
}

func LoadByKey(object interface{}, key interface{}) (interface{}, error) {
    db, connectError := ConnectDatabase()
    defer db.Close()
    if connectError != nil {
        glog.Error(connectError)
        return false, errors.New("无法连接mysql")
    }
    load := db.First(object, key)
    if load.Error != nil {
        glog.Errorf("查询 %T 记录失败, %s", object, load.Error)
        return nil, errors.New("查询失败")
    } else {
        return object, nil
    }

}

