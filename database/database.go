package database

import (
    "log"
    "../config"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "time"
)

type Task struct {
    Task_id         int `gorm:"primary_key:id;AUTO_INCREMENT:id"`
    Task_name       string
    Command         string
    Father_id       []int
    Valid           bool
    Machine_pool    []int
    Owner_id        int
    Schedule_type   string
    Schedule_format string
    Create_time     time.Time `gorm:"-"`
    Update_time     time.Time `gorm:"-"`
}

type Instance struct {
    Instance_id int `gorm:"primary_key:id;AUTO_INCREMENT:id"`
    Task_id     int
    Status      string
    Create_time time.Time `gorm:"-"`
    Update_time time.Time `gorm:"-"`
}

type Log struct {
    Log_id      int `gorm:"primary_key:id;AUTO_INCREMENT:id"`
    Instance_id int
    Machine_id  int
    Output      string
    Create_time time.Time `gorm:"-"`
    Update_time time.Time `gorm:"-"`
}

type Action struct {
    Action_id   int `gorm:"primary_key:id;AUTO_INCREMENT:id"`
    user_id     int
    content     string
    Create_time time.Time `gorm:"-"`
    Update_time time.Time `gorm:"-"`
}

type User struct {
    User_id     int `gorm:"primary_key:id;AUTO_INCREMENT:id"`
    User_name   string
    Password    string
    Email       string
    Create_time time.Time `gorm:"-"`
    Update_time time.Time `gorm:"-"`
}

type Machine struct {
    Machine_id   int `gorm:"primary_key:id;AUTO_INCREMENT:id"`
    Machine_name string
    IP           string
    MAC          string
    CPU_load     int
    Memory_load  int
    alive        bool
    Create_time  time.Time `gorm:"-"`
    Update_time  time.Time `gorm:"-"`
}

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

    log.Print(db.HasTable("users"))
}

func Mock() {

}

func Add() {

}

func Remove() {

}

func Update() {

}

func Query() {

}
