package database

import (
    "testing"
    . "github.com/franela/goblin"
    "github.com/jinzhu/gorm"
    "time"
)

func TestInstance(t *testing.T) {
    g := Goblin(t)

    g.Describe("测试 instance 表插入", func() {
        var db *gorm.DB
        var err error
        var mysqlCount int
        var instances []Instance

        g.Before(func() {
            db, err = ConnectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            db.Exec("truncate table instances;")
            instances = []Instance{
                {TaskID:randomInt(10000), Status:randomString(20)},
                {TaskID:randomInt(10000), Status:randomString(20)},
                {TaskID:randomInt(10000), Status:randomString(20)},
                {TaskID:randomInt(10000), Status:randomString(20)},
                {TaskID:randomInt(10000), Status:randomString(20)},
                {TaskID:randomInt(10000), Status:randomString(20)},
            }

        })

        g.After(func() {
            defer db.Close()
        })

        g.It(" 记录被正确插入", func() {
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

                newInstance, err := new(Instance).LoadByWhere("task_id =?", instance.TaskID)
                g.Assert(err == nil)
                g.Assert(newInstance.Status).Equal(instance.Status)

                newInstance, err = new(Instance).LoadByKey(id + 1)
                g.Assert(err == nil)
                g.Assert(newInstance.TaskID).Equal(instance.TaskID)
                g.Assert(newInstance.Status).Equal(instance.Status)
            }
        })
    })

    g.Describe("测试 instance 表更新", func() {
        var db *gorm.DB
        var err error
        var mysqlCount int
        var instances []Instance

        g.Before(func() {
            db, err = ConnectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            db.Exec("truncate table instances;")
            instances = []Instance{
                {TaskID:randomInt(10000), Status:randomString(20)},
                {TaskID:randomInt(10000), Status:randomString(20)},
                {TaskID:randomInt(10000), Status:randomString(20)},
                {TaskID:randomInt(10000), Status:randomString(20)},
                {TaskID:randomInt(10000), Status:randomString(20)},
                {TaskID:randomInt(10000), Status:randomString(20)},
            }
        })

        g.After(func() {
            defer db.Close()
        })

        g.It("记录被正确更新", func() {
            g.Timeout(10 * time.Second)
            for _, instance := range instances {
                ok, err := instance.DumpToMySQL()
                g.Assert(err == nil).IsTrue()
                g.Assert(ok).IsTrue()

                db.Model(new(Instance)).Where("task_id = ?", instance.TaskID).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)

                oldUpdateTime := instance.UpdatedAt
                oldCreateTime := instance.CreatedAt

                time.Sleep(time.Second)
                newStatus := randomString(3)
                instance.Status = newStatus
                instance.DumpToMySQL()

                newInstance, _ := new(Instance).LoadByWhere("task_id = ?", instance.TaskID)
                g.Assert(newInstance.Status).Equal(newStatus)
                g.Assert(newInstance.UpdatedAt.Sub(oldUpdateTime).Seconds() > 0).IsTrue()
                g.Assert(newInstance.CreatedAt.Format("2006-01-02 15:04:05")).Equal(oldCreateTime.Format("2006-01-02 15:04:05"))
            }

        })
    })

    g.Describe("测试 instance 数据加载", func() {
        var db *gorm.DB
        var err error
        var instances []Instance

        g.Before(func() {
            db, err = ConnectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            db.Exec("truncate table instances;")
            instances = []Instance{
                {TaskID:randomInt(10000), Status:randomString(20)},
                {TaskID:randomInt(10000), Status:randomString(20)},
                {TaskID:randomInt(10000), Status:randomString(20)},
                {TaskID:randomInt(10000), Status:randomString(20)},
                {TaskID:randomInt(10000), Status:randomString(20)},
                {TaskID:randomInt(10000), Status:randomString(20)},
            }

            for _, instance := range instances {
                ok, err := instance.DumpToMySQL()
                g.Assert(ok).IsTrue()
                g.Assert(err == nil).IsTrue()
            }
        })

        g.After(func() {
            defer db.Close()
        })

        g.It("记录通过where条件被正确加载", func() {
            for _, instance := range instances {
                newInstance, err := new(Instance).LoadByWhere("task_id = ?", instance.TaskID)
                g.Assert(err == nil).IsTrue()
                g.Assert(newInstance.Status).Equal(instance.Status)
            }
        })

        g.It("记录通主键被正确加载", func() {
            for id, instance := range instances {
                newInstance, err := new(Instance).LoadByKey(id + 1)
                g.Assert(err == nil).IsTrue()
                g.Assert(newInstance.TaskID).Equal(instance.TaskID)
                g.Assert(newInstance.Status).Equal(instance.Status)
            }
        })

        g.It("记录通过多个where条件被正确加载", func() {
            for id, instance := range instances {
                newInstance, err := new(Instance).LoadByWhere(
                    "id = ? and task_id = ?",
                    id + 1, instance.TaskID,
                )
                g.Assert(err == nil).IsTrue()
                g.Assert(newInstance.Status).Equal(instance.Status)
            }
        })

        g.It("当存在多于一条记录满足where条件时无法实例化，返回异常且对象为nil", func() {
            (&Instance{TaskID:12580, Status:randomString(20)}).DumpToMySQL()
            (&Instance{TaskID:12580, Status:randomString(20)}).DumpToMySQL()
            instance, err := new(Instance).LoadByWhere("task_id = ?", 12580)
            g.Assert(instance == nil).IsTrue()
            g.Assert(err.Error()).Equal("存在多条满足条件的记录，无法实例化")
        })

        g.It("当存在零条记录满足where条件时无法实例化，返回异常且对象为nil", func() {
            newInstance, err := new(Instance).LoadByWhere("task_id = ?", 999999)
            g.Assert(newInstance == nil).IsTrue()
            g.Assert(err.Error()).Equal("不存在满足条件的记录，无法实例化")
        })

    })
}