package database

import (
    "testing"
    . "github.com/franela/goblin"
)

func Test(t *testing.T) {
    g := Goblin(t)
    g.Describe("测试 user.go", func() {
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

}
