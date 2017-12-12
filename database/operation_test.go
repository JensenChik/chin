package database

import (
    "testing"
    . "github.com/franela/goblin"
    "github.com/jinzhu/gorm"
    "time"
)

func TestOperation(t *testing.T) {
    g := Goblin(t)

    g.Describe("测试 operations 表插入", func() {
        var db *gorm.DB
        var err error
        var mysqlCount int
        var operations []Operation

        g.Before(func() {
            db, err = connectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            Truncate("operations")
            operations = []Operation{
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
            }

        })

        g.After(func() {
            defer db.Close()
        })

        g.It("operations记录被正确插入", func() {
            expectedCount := 0
            for id, operation := range operations {
                ok, err := operation.DumpToMySQL()
                g.Assert(ok).IsTrue()
                g.Assert(err == nil).IsTrue()

                expectedCount++
                db.Table("operations").Count(&mysqlCount)
                g.Assert(expectedCount).Equal(mysqlCount)

                db.Model(new(Operation)).Where("id = ?", id + 1).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)

                newOperation, err := new(Operation).LoadByWhere("user_id =?", operation.UserID)
                g.Assert(err == nil)
                g.Assert(newOperation.Content).Equal(operation.Content)

                newOperation, err = new(Operation).LoadByKey(id + 1)
                g.Assert(err == nil)
                g.Assert(newOperation.UserID).Equal(operation.UserID)
                g.Assert(newOperation.Content).Equal(operation.Content)
            }
        })
    })

    g.Describe("测试 opeartions 表更新", func() {
        var db *gorm.DB
        var err error
        var mysqlCount int
        var operations []Operation

        g.Before(func() {
            db, err = connectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            Truncate("operations")
            operations = []Operation{
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
            }
        })

        g.After(func() {
            defer db.Close()
        })

        g.It("opeartions记录被正确更新", func() {
            g.Timeout(10 * time.Second)
            for _, operation := range operations {
                ok, err := operation.DumpToMySQL()
                g.Assert(err == nil).IsTrue()
                g.Assert(ok).IsTrue()

                db.Model(new(Operation)).Where("user_id = ?", operation.UserID).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)

                oldUpdateTime := operation.UpdatedAt
                oldCreateTime := operation.CreatedAt

                time.Sleep(time.Second)
                newContent := randomString(3)
                operation.Content = newContent
                operation.DumpToMySQL()

                newOperation, _ := new(Operation).LoadByWhere("user_id = ?", operation.UserID)
                g.Assert(newOperation.Content).Equal(newContent)
                g.Assert(newOperation.UpdatedAt.Sub(oldUpdateTime).Seconds() > 0).IsTrue()
                g.Assert(newOperation.CreatedAt.Format("2006-01-02 15:04:05")).Equal(oldCreateTime.Format("2006-01-02 15:04:05"))
            }

        })
    })

    g.Describe("测试 operations 数据加载", func() {
        var db *gorm.DB
        var err error
        var operations []Operation

        g.Before(func() {
            db, err = connectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            Truncate("operations")
            operations = []Operation{
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
            }

            for _, operation := range operations {
                ok, err := operation.DumpToMySQL()
                g.Assert(ok).IsTrue()
                g.Assert(err == nil).IsTrue()
            }
        })

        g.After(func() {
            defer db.Close()
        })

        g.It("记录通过where条件被正确加载", func() {
            for _, operation := range operations {
                newOperation, err := new(Operation).LoadByWhere("user_id = ?", operation.UserID)
                g.Assert(err == nil).IsTrue()
                g.Assert(newOperation.Content).Equal(operation.Content)
            }
        })

        g.It("记录通主键被正确加载", func() {
            for id, operation := range operations {
                newOperation, err := new(Operation).LoadByKey(id + 1)
                g.Assert(err == nil).IsTrue()
                g.Assert(newOperation.UserID).Equal(operation.UserID)
                g.Assert(newOperation.Content).Equal(operation.Content)
            }
        })

        g.It("记录通过多个where条件被正确加载", func() {
            for id, operation := range operations {
                newOperation, err := new(Operation).LoadByWhere(
                    "id = ? and user_id = ?",
                    id + 1, operation.UserID,
                )
                g.Assert(err == nil).IsTrue()
                g.Assert(newOperation.Content).Equal(operation.Content)
            }
        })

        g.It("当存在多于一条记录满足where条件时无法实例化，返回异常且对象为nil", func() {
            (&Operation{UserID:12580, Content:randomString(20)}).DumpToMySQL()
            (&Operation{UserID:12580, Content:randomString(20)}).DumpToMySQL()
            operation, err := new(Operation).LoadByWhere("user_id = ?", 12580)
            g.Assert(operation == nil).IsTrue()
            g.Assert(err.Error()).Equal("存在多条满足条件的记录，无法实例化")
        })

        g.It("当存在零条记录满足where条件时无法实例化，返回异常且对象为nil", func() {
            newOperation, err := new(Operation).LoadByWhere("user_id = ?", 999999)
            g.Assert(newOperation == nil).IsTrue()
            g.Assert(err.Error()).Equal("不存在满足条件的记录，无法实例化")
        })

    })
}
