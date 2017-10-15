package database

import (
    "log"
    "../config"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "crypto/md5"
    "encoding/hex"
    "time"
    "math/rand"
)

type Task struct {
    gorm.Model
    Task_name       string
    Command         string
    Father_task     string `gorm:"type:text"`
    Valid           bool
    Machine_pool    string
    Owner_id        int
    Schedule_type   string
    Schedule_format string
}

type Instance struct {
    gorm.Model
    Task_id int `gorm:"index"`
    Status  string
}

type Log struct {
    gorm.Model
    Instance_id int `gorm:"index"`
    Machine_id  int
    Output      string
}

type Action struct {
    gorm.Model
    user_id int
    content string
}

type User struct {
    gorm.Model
    User_name string `gorm:"unique"`
    Password  string
    Email     string
}

type Machine struct {
    gorm.Model
    Machine_name string
    IP           string `gorm:"size:15"`
    MAC          string `gorm:"size:17"`
    CPU_load     int
    Memory_load  int
    alive        bool
}

func to_md5(raw string) string {
    ctx := md5.New()
    ctx.Write([]byte(raw))
    return hex.EncodeToString(ctx.Sum(nil))
}

func random_string(size int) string {
    str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    bytes := []byte(str)
    result := []byte{}
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i := 0; i < size; i++ {
        result = append(result, bytes[r.Intn(len(bytes))])
    }
    return string(result)
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
    db.DropTableIfExists(&Task{}, &Instance{}, &Log{}, &Action{}, &User{}, &Machine{})
    db.CreateTable(&Task{}, &Instance{}, &Log{}, &Action{}, &User{}, &Machine{})
    root_user := User{
        User_name:config.ROOT_NAME,
        Password:to_md5(config.ROOT_PASSWD),
        Email:config.ROOT_MAIL,
    }
    db.Create(&root_user)
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
            Task_name:random_string(16),
            Command:random_string(16),
            Father_task:random_string(32),
            Valid:rand.Float32() < 0.5,
            Machine_pool:random_string(10),
            Owner_id:int(rand.Float32() * 100),
            Schedule_type:random_string(5),
            Schedule_format:random_string(10),
        }
        db.Create(&task)
    }

    log.Print("mock 数据完毕")
}

