package database

import (
    "../config"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "math/rand"
    "github.com/sdbaiguanghe/glog"
)

func ConnectDatabase() (*gorm.DB, error) {
    db, connectError := gorm.Open("mysql", config.SQL_CONN)
    if connectError != nil {
        glog.Error("无法连接mysql, ", connectError)
    }
    return db, connectError
}

func Init() {
    glog.Info("初始化数据库")
    glog.Info("连接mysql")
    db, err := ConnectDatabase()
    if err != nil {
        glog.Fatal("mysql无法链接", err.Error())
    }
    defer db.Close()
    db.DropTableIfExists(&Task{}, &Instance{}, &Log{}, &Action{}, &User{}, &Machine{})
    db.CreateTable(&Task{}, &Instance{}, &Log{}, &Action{}, &User{}, &Machine{})
    rootUser := User{
        UserName:config.ROOT_NAME,
        Password:config.ROOT_PASSWD,
        Email:config.ROOT_MAIL,
    }
    ok, err:= rootUser.DumpToMySQL()
    if ok {
        glog.Info("初始化数据表完毕")
    }
}

func Mock() {
    Init()
    glog.Info("连接 mysql ...")
    glog.Info("开始 mock 数据")
    db, err := gorm.Open("mysql", config.SQL_CONN)
    if err != nil {
        glog.Fatal("连接 mysql 失败: ", err.Error())
    } else {
        glog.Info("连接 mysql 成功")
    }
    defer db.Close()

    glog.Info("开始 mock 表<tasks>")

    for i := 0; i < 100; i++ {
        task := Task{
            TaskName:randomString(16),
            Command:randomString(16),
            FatherTask:randomString(32),
            Valid:rand.Float32() < 0.5,
            MachinePool:randomString(10),
            OwnerID:int(rand.Float32() * 100),
            ScheduleType:randomString(5),
            ScheduleFormat:randomString(10),
        }
        db.Create(&task)
    }

    glog.Info("开始 mock 表<instances>")
    for i := 0; i < 1000; i++ {
        instance := Instance{
            TaskID:int(rand.Float32() * 100),
            Status:randomString(5),
        }
        db.Create(&instance)
    }

    glog.Info("开始 mock 表<logs>")
    for i := 0; i < 10000; i++ {
        log := Log{
            InstanceID:int(rand.Float32() * 1000),
            MachineID:int(rand.Float32() * 10),
            StdOut:randomString(100),
        }
        db.Create(&log)
    }

    glog.Info("开始 mock 表<users>")
    for i := 0; i < 10; i++ {
        user := User{
            UserName:randomString(10),
            Password:randomString(10),
            Email:randomString(5) + "@" + randomString(3) + ".com",
        }
        db.Create(&user)
    }

    glog.Info("开始 mock 表<machines>")
    for i := 0; i < 10; i++ {
        machine := Machine{
            MachineName:randomString(10),
            IP:randomString(10),
            MAC:randomString(10),
            CPULoad:int(rand.Float32() * 100),
            MemoryLoad:int(rand.Float32() * 100),
        }
        db.Create(&machine)
    }

    glog.Info("开始 mock 表<actions>")
    for i := 0; i < 100; i++ {
        action := Action{
            UserID:int(rand.Float32() * 10),
            Content:randomString(20),
        }
        db.Create(&action)
    }

    glog.Info("mock 数据完毕")
}

