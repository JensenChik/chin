package database

import (
    "testing"
    "github.com/jinzhu/gorm"
    . "github.com/franela/goblin"
)

func TestLog(t *testing.T) {
    g := Goblin(t)
    g.Describe("测试 log 表插入", func() {
        var toBeAddLogs []Log
        var db *gorm.DB
        var err error
        var mysqlCount int

        g.Before(func() {
            db, err = ConnectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
        })

        g.BeforeEach(func() {
            Init()
            toBeAddLogs = []Log{
                {InstanceID:1, MachineID:1, Output:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
                {InstanceID:2, MachineID:2, Output:"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"},
                {InstanceID:3, MachineID:3, Output:"cccccccccccccccccccccccccccccc"},
                {InstanceID:4, MachineID:4, Output:"dddddddddddddddddddddddddddddd"},
                {InstanceID:5, MachineID:5, Output:"eeeeeeeeeeeeeeeeeeeeeeeeeeeeee"},
                {InstanceID:6, MachineID:6, Output:"ffffffffffffffffffffffffffffff"},
                {InstanceID:7, MachineID:7, Output:"gggggggggggggggggggggggggggggg"},
            }
        })

        g.After(func() {
            defer db.Close()
        })

        g.It("log 记录正确被插入", func() {
            expectedCount := 0
            var log_copy Log
            for _, log := range toBeAddLogs {
                log_copy = log
                ok := log_copy.DumpToMySQL()
                g.Assert(ok).IsTrue()
                expectedCount++
                db.Table("logs").Count(&mysqlCount)
                g.Assert(expectedCount).Equal(mysqlCount)
                db.Where("id = ?", log_copy.ID).Find(&log_copy).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)
                db.Where("id = ?", log_copy.ID).First(&log_copy)
                g.Assert(log.Output).Equal(log_copy.Output)
            }
        })

    })
}
