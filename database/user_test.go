package database

import (
    "testing"
    . "github.com/franela/goblin"
    "github.com/jinzhu/gorm"
    "time"
)

func TestUser(t *testing.T) {
    g := Goblin(t)

    g.Describe("测试 ExistsUser@user.go", func() {
        var db *gorm.DB
        var err error

        g.Before(func() {
            db, err = ConnectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            db.Exec("truncate table users;")
            ok, _ := (&User{UserName:"chin", Password:"root"}).DumpToMySQL()
            g.Assert(ok).IsTrue()
        })
        g.After(func() {
            db.Exec("truncate table users;")
            defer db.Close()
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
            toBeAddUsers = []User{
                {UserName:"A", Password:"1", Email:"A@1"},
                {UserName:"B", Password:"2", Email:"B@2"},
                {UserName:"C", Password:"3", Email:"C@3"},
                {UserName:"D", Password:"4", Email:"C@3"},
            }
            db.Exec("truncate table users;")
        })

        g.After(func() {
            db.Exec("truncate table users;")
            defer db.Close()
        })

        g.It(" user_name 不存在的记录正确被插入", func() {
            expectedCount := 0
            for _, user := range toBeAddUsers {
                ok, err := user.DumpToMySQL()
                g.Assert(ok).IsTrue()
                g.Assert(err == nil).IsTrue()

                expectedCount++
                db.Table("users").Count(&mysqlCount)
                g.Assert(expectedCount).Equal(mysqlCount)

                db.Model(new(User)).Where("user_name = ?", user.UserName).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)

                db.Model(new(User)).Where("email = ?", user.Email).Count(&mysqlCount)
                switch user.UserName {
                case "A", "B", "C":
                    g.Assert(mysqlCount).Equal(1)
                case "D":
                    g.Assert(mysqlCount).Equal(2)
                default:
                    g.Fail("非法用户名")
                }

                newUser, err := new(User).LoadByWhere("user_name =?", user.UserName)
                g.Assert(err == nil)
                g.Assert(newUser.Password).Equal(toMD5(user.Password))
            }
        })

        g.It("插入重复的user_name记录返回失败", func() {
            for _, user := range toBeAddUsers {
                ok, err := user.DumpToMySQL()
                g.Assert(err == nil).IsTrue()
                g.Assert(ok).IsTrue()

                db.Model(new(User)).Where("user_name = ?", user.UserName).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)

                randomPasswd := randomString(10)
                randomEmail := randomString(10)
                ok, err = (&User{
                    UserName: user.UserName,
                    Password: randomPasswd,
                    Email: randomEmail,
                }).DumpToMySQL()
                g.Assert(ok).IsFalse()
                g.Assert(err != nil).IsTrue()

                db.Model(new(User)).Where("user_name = ?", user.UserName).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)

                newUser, err := new(User).LoadByWhere("user_name = ?", user.UserName)
                g.Assert(newUser.Password != toMD5(randomPasswd)).IsTrue()
                g.Assert(newUser.Password).Equal(toMD5(user.Password))
                g.Assert(newUser.Email != randomEmail).IsTrue()
                g.Assert(newUser.Email).Equal(user.Email)
            }
        })
    })

    g.Describe("测试user表更新", func() {
        var users []User
        var db *gorm.DB
        var err error
        var mysqlCount int

        g.Before(func() {
            db, err = ConnectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            db.Exec("truncate table users;")
            users = []User{
                {UserName:"A", Password:"1", Email:"A@1"},
                {UserName:"B", Password:"2", Email:"B@2"},
                {UserName:"C", Password:"3", Email:"C@3"},
                {UserName:"D", Password:"4", Email:"D@4"},
            }
        })

        g.After(func() {
            db.Exec("truncate table users;")
            defer db.Close()
        })

        g.Xit("记录被正确更新", func() {
            g.Timeout(10 * time.Second)
            for _, user := range users {
                ok, err := user.DumpToMySQL()
                g.Assert(err == nil).IsTrue()
                g.Assert(ok).IsTrue()

                db.Model(new(User)).Where("user_name = ?", user.UserName).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)

                oldUpdateTime := user.UpdatedAt
                oldCreateTime := user.CreatedAt

                time.Sleep(2 * time.Second)
                newPasswd := randomString(30)
                newEmail := randomString(30)
                user.Password = newPasswd
                user.Email = newEmail
                user.DumpToMySQL()

                newUser, _ := new(User).LoadByWhere("user_name = ?", user.UserName)
                g.Assert(newUser.Password).Equal(toMD5(newPasswd))
                g.Assert(newUser.Email).Equal(newEmail)
                g.Assert(newUser.UpdatedAt.Sub(oldUpdateTime).Seconds() > 0).IsTrue()
                g.Assert(newUser.CreatedAt.Format("2006-01-02 15:04:05")).Equal(oldCreateTime.Format("2006-01-02 15:04:05"))
            }

        })
    })

    g.Describe("测试 user 数据加载", func() {
        var db *gorm.DB
        var err error
        var users []User

        g.Before(func() {
            db, err = ConnectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            db.Exec("truncate table users;")
            users = []User{
                {UserName:"A", Password:"1", Email:"A@1"},
                {UserName:"B", Password:"2", Email:"B@2"},
                {UserName:"C", Password:"3", Email:"C@3"},
                {UserName:"D", Password:"4", Email:"C@3"},
            }

            for _, user := range users {
                ok, err := user.DumpToMySQL()
                g.Assert(ok).IsTrue()
                g.Assert(err == nil).IsTrue()
            }
        })

        g.After(func() {
            db.Exec("truncate table users;")
            defer db.Close()
        })

        g.It("记录通过where条件被正确加载", func() {
            for _, user := range users {
                newUser, err := new(User).LoadByWhere("user_name = ?", user.UserName)
                g.Assert(err == nil).IsTrue()
                g.Assert(newUser.UserName).Equal(user.UserName)
                g.Assert(newUser.Password).Equal(toMD5(user.Password))
                g.Assert(newUser.Email).Equal(user.Email)
            }
        })

        g.It("记录通主键被正确加载", func() {
            for id, user := range users {
                newUser, err := new(User).LoadByKey(id + 1)
                g.Assert(err == nil).IsTrue()
                g.Assert(newUser.UserName).Equal(user.UserName)
                g.Assert(newUser.Password).Equal(toMD5(user.Password))
                g.Assert(newUser.Email).Equal(user.Email)
            }
        })

        g.It("记录通过多个where条件被正确加载", func() {
            for id, user := range users {
                newUser, err := new(User).LoadByWhere(
                    "id = ? and user_name = ? and email = ?",
                    id + 1, user.UserName, user.Email,
                )
                g.Assert(err == nil).IsTrue()
                g.Assert(newUser.UserName).Equal(user.UserName)
                g.Assert(newUser.Password).Equal(toMD5(user.Password))
                g.Assert(newUser.Email).Equal(user.Email)
            }
        })

        g.It("当存在多于一条记录满足where条件时无法实例化，返回异常且对象为nil", func() {
            newUser, err := new(User).LoadByWhere("email = ?", users[2].Email)
            g.Assert(newUser == nil).IsTrue()
            g.Assert(err.Error()).Equal("存在多条满足条件的记录，无法实例化")
        })

        g.It("当存在零条记录满足where条件时无法实例化，返回异常且对象为nil", func() {
            tmpUser, err := new(User).LoadByWhere("email = ?", "fuck@shit")
            g.Assert(tmpUser == nil).IsTrue()
            g.Assert(err.Error()).Equal("不存在满足条件的记录，无法实例化")
        })

    })
}

