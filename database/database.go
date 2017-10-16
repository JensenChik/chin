package database

import (
    "log"
    "../config"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "math/rand"
)

func Init() {

    log.Print("初始化数据库")
    log.Print("连接 mysql ...")
    db, err := gorm.Open("mysql", config.SQL_CONN)
    if err != nil {
        log.Fatal("连接 mysql 失败: ", err.Error())
    } else {
        log.Print("连接 mysql 成功")
    }
    defer db.Close()
    db.DropTableIfExists(&Task{}, &Instance{}, &Log{}, &Action{}, &User{}, &Machine{})
    db.CreateTable(&Task{}, &Instance{}, &Log{}, &Action{}, &User{}, &Machine{})
    rootUser := User{
        UserName:config.ROOT_NAME,
        Password:toMD5(config.ROOT_PASSWD),
        Email:config.ROOT_MAIL,
    }
    db.Create(&rootUser)
    log.Print("初始化数据表完毕")
}

func Mock() {
    Init()
    log.Print("开始mock 数据...")
    log.Print("连接 mysql ...")
    db, err := gorm.Open("mysql", config.SQL_CONN)
    if err != nil {
        log.Fatal("连接 mysql 失败: ", err.Error())
    } else {
        log.Print("连接 mysql 成功")
    }
    defer db.Close()

    log.Print("开始 mock 表<tasks>")
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

    log.Print("开始 mock 表<instances>")
    for i := 0; i < 1000; i++ {
        instance := Instance{
            TaskID:int(rand.Float32() * 100),
            Status:randomString(5),
        }
        db.Create(&instance)
    }

    log.Print("开始 mock 表<logs>")
    for i := 0; i < 10000; i++ {
        log := Log{
            InstanceID:int(rand.Float32() * 1000),
            MachineID:int(rand.Float32() * 10),
            Output:randomString(100),
        }
        db.Create(&log)
    }

    log.Print("开始 mock 表<users>")
    for i := 0; i < 10; i++ {
        user := User{
            UserName:randomString(10),
            Password:toMD5(randomString(10)),
            Email:randomString(5) + "@" + randomString(3) + ".com",
        }
        db.Create(&user)
    }

    log.Print("开始 mock 表<machines>")
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

    log.Print("开始 mock 表<actions>")
    for i := 0; i < 100; i++ {
        action := Action{
            UserID:int(rand.Float32() * 10),
            Content:randomString(20),
        }
        db.Create(&action)
    }

    log.Print("mock 数据完毕")
}

