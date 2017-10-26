package database

import (
    "testing"
    . "github.com/franela/goblin"
    "github.com/jinzhu/gorm"
    "time"
)

func TestLog(t *testing.T) {
    g := Goblin(t)

    g.Describe("测试 log 表插入", func() {
        var db *gorm.DB
        var err error
        var mysqlCount int
        var logs []Log

        g.Before(func() {
            db, err = ConnectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            db.Exec("truncate table logs;")
            logs = []Log{
                {InstanceID:randomInt(10000), MachineID:randomInt(1000), StdOut:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
                {InstanceID:randomInt(10000), MachineID:randomInt(1000), StdOut:"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"},
                {InstanceID:randomInt(10000), MachineID:randomInt(1000), StdOut:"ccccccccccccccccccccccccccccccccccc"},
                {InstanceID:randomInt(10000), MachineID:randomInt(1000), StdOut:"ddddddddddddddddddddddddddddddddddd"},
                {InstanceID:randomInt(10000), MachineID:randomInt(1000), StdOut:"eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"},
                {InstanceID:randomInt(10000), MachineID:randomInt(1000), StdOut:"fffffffffffffffffffffffffffffffffff"},
            }

        })

        g.After(func() {
            defer db.Close()
        })

        g.It(" 记录被正确插入", func() {
            expectedCount := 0
            for id, log := range logs {
                ok, err := log.DumpToMySQL()
                g.Assert(ok).IsTrue()
                g.Assert(err == nil).IsTrue()

                expectedCount++
                db.Table("logs").Count(&mysqlCount)
                g.Assert(expectedCount).Equal(mysqlCount)

                db.Model(new(Log)).Where("id = ?", id + 1).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)

                newLog, err := new(Log).LoadByWhere("id =?", id + 1)
                g.Assert(err == nil)
                g.Assert(newLog.InstanceID).Equal(log.InstanceID)
                g.Assert(newLog.MachineID).Equal(log.MachineID)
                g.Assert(newLog.StdOut).Equal(log.StdOut)

                newLog, err = new(Log).LoadByKey(id + 1)
                g.Assert(err == nil)
                g.Assert(newLog.InstanceID).Equal(log.InstanceID)
                g.Assert(newLog.MachineID).Equal(log.MachineID)
                g.Assert(newLog.StdOut).Equal(log.StdOut)
            }
        })
    })

    g.Describe("测试 log 表更新", func() {
        var db *gorm.DB
        var err error
        var mysqlCount int
        var logs []Log

        g.Before(func() {
            db, err = ConnectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            db.Exec("truncate table logs;")
            logs = []Log{
                {InstanceID:randomInt(10000), MachineID:randomInt(1000), StdOut:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
                {InstanceID:randomInt(10000), MachineID:randomInt(1000), StdOut:"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"},
                {InstanceID:randomInt(10000), MachineID:randomInt(1000), StdOut:"ccccccccccccccccccccccccccccccccccc"},
                {InstanceID:randomInt(10000), MachineID:randomInt(1000), StdOut:"ddddddddddddddddddddddddddddddddddd"},
                {InstanceID:randomInt(10000), MachineID:randomInt(1000), StdOut:"eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"},
                {InstanceID:randomInt(10000), MachineID:randomInt(1000), StdOut:"fffffffffffffffffffffffffffffffffff"},
            }
        })

        g.After(func() {
            defer db.Close()
        })

        g.It("记录被正确更新", func() {
            g.Timeout(10 * time.Second)
            for id, log := range logs {
                ok, err := log.DumpToMySQL()
                g.Assert(err == nil).IsTrue()
                g.Assert(ok).IsTrue()

                db.Model(new(Log)).Where("id = ?", id + 1).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)

                oldUpdateTime := log.UpdatedAt
                oldCreateTime := log.CreatedAt

                time.Sleep(time.Second)
                newInstanceID := randomInt(1000)
                newMachineID := randomInt(1000)
                newStdOut := randomString(100)
                log.InstanceID = newInstanceID
                log.MachineID = newMachineID
                log.StdOut = newStdOut
                log.DumpToMySQL()

                newLog, _ := new(Log).LoadByWhere("id = ?", id + 1)
                g.Assert(newLog.InstanceID).Equal(newInstanceID)
                g.Assert(newLog.MachineID).Equal(newMachineID)
                g.Assert(newLog.StdOut).Equal(newStdOut)
                g.Assert(newLog.UpdatedAt.Sub(oldUpdateTime).Seconds() > 0).IsTrue()
                g.Assert(newLog.CreatedAt.Format("2006-01-02 15:04:05")).Equal(oldCreateTime.Format("2006-01-02 15:04:05"))
            }

        })
    })

    g.Describe("测试 log 数据加载", func() {
        var db *gorm.DB
        var err error
        var logs []Log

        g.Before(func() {
            db, err = ConnectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            db.Exec("truncate table logs;")
            logs = []Log{
                {InstanceID:randomInt(10000), MachineID:randomInt(1000), StdOut:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
                {InstanceID:randomInt(10000), MachineID:randomInt(1000), StdOut:"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"},
                {InstanceID:randomInt(10000), MachineID:randomInt(1000), StdOut:"ccccccccccccccccccccccccccccccccccc"},
                {InstanceID:randomInt(10000), MachineID:randomInt(1000), StdOut:"ddddddddddddddddddddddddddddddddddd"},
                {InstanceID:randomInt(10000), MachineID:randomInt(1000), StdOut:"eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"},
                {InstanceID:randomInt(10000), MachineID:randomInt(1000), StdOut:"fffffffffffffffffffffffffffffffffff"},
            }

            for _, log := range logs {
                ok, err := log.DumpToMySQL()
                g.Assert(ok).IsTrue()
                g.Assert(err == nil).IsTrue()
            }
        })

        g.After(func() {
            defer db.Close()
        })

        g.It("记录通过where条件被正确加载", func() {
            for id, log := range logs {
                newLog, err := new(Log).LoadByWhere("id = ?", id + 1)
                g.Assert(err == nil).IsTrue()
                g.Assert(newLog.InstanceID).Equal(log.InstanceID)
                g.Assert(newLog.MachineID).Equal(log.MachineID)
                g.Assert(newLog.StdOut).Equal(log.StdOut)
            }
        })

        g.It("记录通主键被正确加载", func() {
            for id, log := range logs {
                newLog, err := new(Log).LoadByKey(id + 1)
                g.Assert(err == nil).IsTrue()
                g.Assert(newLog.InstanceID).Equal(log.InstanceID)
                g.Assert(newLog.MachineID).Equal(log.MachineID)
                g.Assert(newLog.StdOut).Equal(log.StdOut)
            }
        })

        g.It("记录通过多个where条件被正确加载", func() {
            for id, log := range logs {
                newLog, err := new(Log).LoadByWhere(
                    "id = ? and instance_id = ? and machine_id = ?",
                    id + 1, log.InstanceID, log.MachineID,
                )
                g.Assert(err == nil).IsTrue()
                g.Assert(newLog.InstanceID).Equal(log.InstanceID)
                g.Assert(newLog.MachineID).Equal(log.MachineID)
                g.Assert(newLog.StdOut).Equal(log.StdOut)
            }
        })

        g.It("当存在多于一条记录满足where条件时无法实例化，返回异常且对象为nil", func() {
            (&Log{InstanceID:12580}).DumpToMySQL()
            (&Log{InstanceID:12580}).DumpToMySQL()
            log, err := new(Log).LoadByWhere("instance_id = ?", 12580)
            g.Assert(log == nil).IsTrue()
            g.Assert(err.Error()).Equal("存在多条满足条件的记录，无法实例化")
        })

        g.It("当存在零条记录满足where条件时无法实例化，返回异常且对象为nil", func() {
            newLog, err := new(Log).LoadByWhere("instance_id = ?", 999999)
            g.Assert(newLog == nil).IsTrue()
            g.Assert(err.Error()).Equal("不存在满足条件的记录，无法实例化")
        })

    })
}
