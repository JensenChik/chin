package model

import (
    "testing"
    . "github.com/franela/goblin"
    "github.com/jinzhu/gorm"
    "time"
    "../tools/random"
)

func TestInstance(t *testing.T) {
    g := Goblin(t)

    g.Describe("测试 instances 表插入", func() {
        var db *gorm.DB
        var err error
        var mysqlCount int
        var instances []Instance

        g.Before(func() {
            Truncate("instances")
            db, err = connectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            instances = []Instance{
                {JobID:random.Int(10000), MachineID:random.Int(1000), StdOut:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
                {JobID:random.Int(10000), MachineID:random.Int(1000), StdOut:"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"},
                {JobID:random.Int(10000), MachineID:random.Int(1000), StdOut:"ccccccccccccccccccccccccccccccccccc"},
                {JobID:random.Int(10000), MachineID:random.Int(1000), StdOut:"ddddddddddddddddddddddddddddddddddd"},
                {JobID:random.Int(10000), MachineID:random.Int(1000), StdOut:"eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"},
                {JobID:random.Int(10000), MachineID:random.Int(1000), StdOut:"fffffffffffffffffffffffffffffffffff"},
            }

        })

        g.After(func() {
            defer db.Close()
        })

        g.It("instance记录被正确插入", func() {
            expectedCount := 0
            for id, instance := range instances {
                ok, err := instance.DumpToMySQL()
                g.Assert(ok).IsTrue()
                g.Assert(err == nil).IsTrue()

                expectedCount++
                db.Table("instances").Count(&mysqlCount)
                g.Assert(expectedCount).Equal(mysqlCount)

                db.Model(new(Instance)).Where("id = ?", id + 1).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)

                newInstance, err := new(Instance).LoadByWhere("id =?", id + 1)
                g.Assert(err == nil)
                g.Assert(newInstance.JobID).Equal(instance.JobID)
                g.Assert(newInstance.MachineID).Equal(instance.MachineID)
                g.Assert(newInstance.StdOut).Equal(instance.StdOut)

                newInstance, err = new(Instance).LoadByKey(id + 1)
                g.Assert(err == nil)
                g.Assert(newInstance.JobID).Equal(instance.JobID)
                g.Assert(newInstance.MachineID).Equal(instance.MachineID)
                g.Assert(newInstance.StdOut).Equal(instance.StdOut)
            }
        })
    })

    g.Describe("测试 instances 表更新", func() {
        var db *gorm.DB
        var err error
        var mysqlCount int
        var instances []Instance

        g.Before(func() {
            db, err = connectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            Truncate("instances")
            instances = []Instance{
                {JobID:random.Int(10000), MachineID:random.Int(1000), StdOut:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
                {JobID:random.Int(10000), MachineID:random.Int(1000), StdOut:"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"},
                {JobID:random.Int(10000), MachineID:random.Int(1000), StdOut:"ccccccccccccccccccccccccccccccccccc"},
                {JobID:random.Int(10000), MachineID:random.Int(1000), StdOut:"ddddddddddddddddddddddddddddddddddd"},
                {JobID:random.Int(10000), MachineID:random.Int(1000), StdOut:"eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"},
                {JobID:random.Int(10000), MachineID:random.Int(1000), StdOut:"fffffffffffffffffffffffffffffffffff"},
            }
        })

        g.After(func() {
            defer db.Close()
        })

        g.It("instances记录被正确更新", func() {
            g.Timeout(10 * time.Second)
            for id, instance := range instances {
                ok, err := instance.DumpToMySQL()
                g.Assert(err == nil).IsTrue()
                g.Assert(ok).IsTrue()

                db.Model(new(Instance)).Where("id = ?", id + 1).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)

                oldUpdateTime := instance.UpdatedAt
                oldCreateTime := instance.CreatedAt

                time.Sleep(time.Second)
                newInstanceID := random.Int(1000)
                newMachineID := random.Int(1000)
                newStdOut := random.String(100)
                instance.JobID = newInstanceID
                instance.MachineID = newMachineID
                instance.StdOut = newStdOut
                instance.DumpToMySQL()

                newInstance, _ := new(Instance).LoadByWhere("id = ?", id + 1)
                g.Assert(newInstance.JobID).Equal(newInstanceID)
                g.Assert(newInstance.MachineID).Equal(newMachineID)
                g.Assert(newInstance.StdOut).Equal(newStdOut)
                g.Assert(newInstance.UpdatedAt.Sub(oldUpdateTime).Seconds() > 0).IsTrue()
                g.Assert(newInstance.CreatedAt.Format("2006-01-02 15:04:05")).Equal(oldCreateTime.Format("2006-01-02 15:04:05"))
            }

        })
    })

    g.Describe("测试 instances 数据加载", func() {
        var instances []Instance

        g.Before(func() {
            Truncate("instances")
            instances = []Instance{
                {JobID:random.Int(10000), MachineID:random.Int(1000), StdOut:"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
                {JobID:random.Int(10000), MachineID:random.Int(1000), StdOut:"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"},
                {JobID:random.Int(10000), MachineID:random.Int(1000), StdOut:"ccccccccccccccccccccccccccccccccccc"},
                {JobID:random.Int(10000), MachineID:random.Int(1000), StdOut:"ddddddddddddddddddddddddddddddddddd"},
                {JobID:random.Int(10000), MachineID:random.Int(1000), StdOut:"eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"},
                {JobID:random.Int(10000), MachineID:random.Int(1000), StdOut:"fffffffffffffffffffffffffffffffffff"},
            }

            for _, instance := range instances {
                ok, err := instance.DumpToMySQL()
                g.Assert(ok).IsTrue()
                g.Assert(err == nil).IsTrue()
            }
        })

        g.It("记录通过where条件被正确加载", func() {
            for id, instance := range instances {
                newInstance, err := new(Instance).LoadByWhere("id = ?", id + 1)
                g.Assert(err == nil).IsTrue()
                g.Assert(newInstance.JobID).Equal(instance.JobID)
                g.Assert(newInstance.MachineID).Equal(instance.MachineID)
                g.Assert(newInstance.StdOut).Equal(instance.StdOut)
            }
        })

        g.It("记录通主键被正确加载", func() {
            for id, instance := range instances {
                newInstance, err := new(Instance).LoadByKey(id + 1)
                g.Assert(err == nil).IsTrue()
                g.Assert(newInstance.JobID).Equal(instance.JobID)
                g.Assert(newInstance.MachineID).Equal(instance.MachineID)
                g.Assert(newInstance.StdOut).Equal(instance.StdOut)
            }
        })

        g.It("记录通过多个where条件被正确加载", func() {
            for id, instance := range instances {
                newInstance, err := new(Instance).LoadByWhere(
                    "id = ? and job_id = ? and machine_id = ?",
                    id + 1, instance.JobID, instance.MachineID,
                )
                g.Assert(err == nil).IsTrue()
                g.Assert(newInstance.JobID).Equal(instance.JobID)
                g.Assert(newInstance.MachineID).Equal(instance.MachineID)
                g.Assert(newInstance.StdOut).Equal(instance.StdOut)
            }
        })

        g.It("当存在多于一条记录满足where条件时无法实例化，返回异常且对象为nil", func() {
            (&Instance{JobID:12580}).DumpToMySQL()
            (&Instance{JobID:12580}).DumpToMySQL()
            instance, err := new(Instance).LoadByWhere("job_id = ?", 12580)
            g.Assert(instance == nil).IsTrue()
            g.Assert(err.Error()).Equal("存在多条满足条件的记录，无法实例化")
        })

        g.It("当存在零条记录满足where条件时无法实例化，返回异常且对象为nil", func() {
            newInstance, err := new(Instance).LoadByWhere("job_id = ?", 999999)
            g.Assert(newInstance == nil).IsTrue()
            g.Assert(err.Error()).Equal("不存在满足条件的记录，无法实例化")
        })

    })
}
