package database

import (
    "testing"
    . "github.com/franela/goblin"
    "github.com/jinzhu/gorm"
    "time"
)

func TestAction(t *testing.T) {
    g := Goblin(t)

    g.Describe("测试 action 表插入", func() {
        var db *gorm.DB
        var err error
        var mysqlCount int
        var actions []Action

        g.Before(func() {
            db, err = ConnectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            db.Exec("truncate table actions;")
            actions = []Action{
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

        g.It(" 记录被正确插入", func() {
            expectedCount := 0
            for id, action := range actions {
                ok, err := action.DumpToMySQL()
                g.Assert(ok).IsTrue()
                g.Assert(err == nil).IsTrue()

                expectedCount++
                db.Table("actions").Count(&mysqlCount)
                g.Assert(expectedCount).Equal(mysqlCount)

                db.Model(new(Action)).Where("id = ?", id + 1).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)

                newAction, err := new(Action).LoadByWhere("user_id =?", action.UserID)
                g.Assert(err == nil)
                g.Assert(newAction.Content).Equal(action.Content)

                newAction, err = new(Action).LoadByKey(id + 1)
                g.Assert(err == nil)
                g.Assert(newAction.UserID).Equal(action.UserID)
                g.Assert(newAction.Content).Equal(action.Content)
            }
        })
    })

    g.Describe("测试 action 表更新", func() {
        var db *gorm.DB
        var err error
        var mysqlCount int
        var actions []Action

        g.Before(func() {
            db, err = ConnectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            db.Exec("truncate table actions;")
            actions = []Action{
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

        g.It("记录被正确更新", func() {
            g.Timeout(10 * time.Second)
            for _, action := range actions {
                ok, err := action.DumpToMySQL()
                g.Assert(err == nil).IsTrue()
                g.Assert(ok).IsTrue()

                db.Model(new(Action)).Where("user_id = ?", action.UserID).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)

                oldUpdateTime := action.UpdatedAt
                oldCreateTime := action.CreatedAt

                time.Sleep(time.Second)
                newContent := randomString(3)
                action.Content = newContent
                action.DumpToMySQL()

                newAction, _ := new(Action).LoadByWhere("user_id = ?", action.UserID)
                g.Assert(newAction.Content).Equal(newContent)
                g.Assert(newAction.UpdatedAt.Sub(oldUpdateTime).Seconds() > 0).IsTrue()
                g.Assert(newAction.CreatedAt.Format("2006-01-02 15:04:05")).Equal(oldCreateTime.Format("2006-01-02 15:04:05"))
            }

        })
    })

    g.Describe("测试 action 数据加载", func() {
        var db *gorm.DB
        var err error
        var actions []Action

        g.Before(func() {
            db, err = ConnectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            db.Exec("truncate table actions;")
            actions = []Action{
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
                {UserID:randomInt(10000), Content:randomString(20)},
            }

            for _, action := range actions {
                ok, err := action.DumpToMySQL()
                g.Assert(ok).IsTrue()
                g.Assert(err == nil).IsTrue()
            }
        })

        g.After(func() {
            defer db.Close()
        })

        g.It("记录通过where条件被正确加载", func() {
            for _, action := range actions {
                newAction, err := new(Action).LoadByWhere("user_id = ?", action.UserID)
                g.Assert(err == nil).IsTrue()
                g.Assert(newAction.Content).Equal(action.Content)
            }
        })

        g.It("记录通主键被正确加载", func() {
            for id, action := range actions {
                newAction, err := new(Action).LoadByKey(id + 1)
                g.Assert(err == nil).IsTrue()
                g.Assert(newAction.UserID).Equal(action.UserID)
                g.Assert(newAction.Content).Equal(action.Content)
            }
        })

        g.It("记录通过多个where条件被正确加载", func() {
            for id, action := range actions {
                newAction, err := new(Action).LoadByWhere(
                    "id = ? and user_id = ?",
                    id + 1, action.UserID,
                )
                g.Assert(err == nil).IsTrue()
                g.Assert(newAction.Content).Equal(action.Content)
            }
        })

        g.It("当存在多于一条记录满足where条件时无法实例化，返回异常且对象为nil", func() {
            (&Action{UserID:12580, Content:randomString(20)}).DumpToMySQL()
            (&Action{UserID:12580, Content:randomString(20)}).DumpToMySQL()
            action, err := new(Action).LoadByWhere("user_id = ?", 12580)
            g.Assert(action == nil).IsTrue()
            g.Assert(err.Error()).Equal("存在多条满足条件的记录，无法实例化")
        })

        g.It("当存在零条记录满足where条件时无法实例化，返回异常且对象为nil", func() {
            newAction, err := new(Action).LoadByWhere("user_id = ?", 999999)
            g.Assert(newAction == nil).IsTrue()
            g.Assert(err.Error()).Equal("不存在满足条件的记录，无法实例化")
        })

    })
}
