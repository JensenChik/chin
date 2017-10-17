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

        g.It("当存在该用户时应当返回true")
        g.It("当不存在该用户时应当返回false")

    })

}
