package database

import (
    "testing"
    . "github.com/franela/goblin"
    "github.com/jinzhu/gorm"
)

func TestUser(t *testing.T) {
    g := Goblin(t)
    g.Describe("测试 ExistsUser@user.go", func() {
        g.Before(func() {
            Init()
        })
        g.After(func() {
            Init()
        })

        g.It("当存在该用户且密码匹配时返回true", func() {
            g.Assert(ExistsUser("chin", "root")).IsTrue()
        })
        g.It("当存在该用户,但密码不匹配时返回false", func() {
            g.Assert(ExistsUser("chin", "root+1s")).IsFalse()
        })
        g.It("当不存在该用户时应当返回false", func() {
            g.Assert(ExistsUser("hahaha", "+1s")).IsFalse()
        })

    })

    g.Describe("测试user表插入", func() {
        var toBeAddUsers []User
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
            toBeAddUsers = []User{
                {UserName:"A", Password:"1", Email:"A@1"},
                {UserName:"B", Password:"2", Email:"B@2"},
                {UserName:"C", Password:"3", Email:"C@3"},
                {UserName:"D", Password:"4", Email:"D@4"},
            }
        })

        g.After(func() {
            defer db.Close()
        })

        g.It(" user_name 不存在的记录正确被插入", func() {
            expectedCount := 1 //init的时候会插入一条root记录
            var user_copy User
            for _, user := range toBeAddUsers {
                user_copy = user
                ok := user_copy.DumpToMySQL()
                g.Assert(ok).IsTrue()
                expectedCount++
                db.Table("users").Count(&mysqlCount)
                g.Assert(expectedCount).Equal(mysqlCount)
                db.Where("user_name = ?", user_copy.UserName).Find(&user_copy).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)
                db.Where("user_name = ?", user_copy.UserName).First(&user_copy)
                g.Assert(user_copy.Password).Eql(toMD5(user.Password))
            }
        })

        g.It("插入重复的user_name记录返回失败", func() {
            for _, user := range toBeAddUsers {
                ok := user.DumpToMySQL()
                g.Assert(ok).IsTrue()
                db.Where("user_name = ?", user.UserName).Find(&user).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)
                anotherUser := User{
                    UserName:user.UserName,
                    Password:randomString(3),
                    Email:randomString(10),
                }
                ok = anotherUser.DumpToMySQL()
                g.Assert(ok).IsFalse()
                db.Where("user_name = ?", user.UserName).Find(&user).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)
            }
        })
    })
}
