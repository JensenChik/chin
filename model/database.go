package model

import (
    "math/rand"

    "../config"
    "../tools/random"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/sdbaiguanghe/glog"
)

func connectDatabase() (*gorm.DB, error) {
    db, connectError := gorm.Open("mysql", config.SQL_CONN)
    if connectError != nil {
        glog.Error("无法连接mysql, ", connectError)
    }
    return db, connectError
}

func Init() {
    glog.Info("初始化数据库")
    glog.Info("连接mysql")
    db, err := connectDatabase()
    if err != nil {
        glog.Fatal("mysql无法链接", err.Error())
    }
    defer db.Close()
    db.DropTableIfExists(&Task{}, &Job{}, &Instance{}, &Operation{}, &User{}, &Machine{})
    db.CreateTable(&Task{}, &Job{}, &Instance{}, &Operation{}, &User{}, &Machine{})
    rootUser := User{
        UserName: config.ROOT_NAME,
        Password: config.ROOT_PASSWD,
        Email:    config.ROOT_MAIL,
    }
    ok, err := rootUser.DumpToMySQL()
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
            TaskName:    random.String(16),
            Command:     random.String(16),
            FatherTask:  random.String(32),
            Valid:       rand.Float32() < 0.5,
            MachinePool: random.String(10),
            OwnerID:     random.Int(100),
        }
        db.Create(&task)
    }

    glog.Info("开始 mock 表<jobs>")
    for i := 0; i < 1000; i++ {
        job := Job{
            TaskID: random.Int(100),
            Status: random.String(5),
        }
        db.Create(&job)
    }

    glog.Info("开始 mock 表<instances>")
    for i := 0; i < 10000; i++ {
        instance := Instance{
            JobID:     random.Int(1000),
            MachineID: random.Int(10),
            StdOut:    random.String(100),
        }
        db.Create(&instance)
    }

    glog.Info("开始 mock 表<users>")
    for i := 0; i < 10; i++ {
        user := User{
            UserName: random.String(10),
            Password: random.String(10),
            Email:    random.String(5) + "@" + random.String(3) + ".com",
        }
        db.Create(&user)
    }

    glog.Info("开始 mock 表<machines>")
    for i := 0; i < 10; i++ {
        machine := Machine{
            MachineName: random.String(10),
            IP:          random.String(10),
            MAC:         random.String(10),
            CPULoad:     random.Float(),
            MemoryLoad:  random.Float(),
        }
        db.Create(&machine)
    }

    glog.Info("开始 mock 表<operations>")
    for i := 0; i < 100; i++ {
        operation := Operation{
            UserID:  random.Int(10),
            Content: random.String(20),
        }
        db.Create(&operation)
    }

    glog.Info("mock 数据完毕")
}
