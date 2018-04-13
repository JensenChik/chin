package secure

import (
    "testing"
    . "github.com/franela/goblin"
)

func TestSecure(t *testing.T) {
    g := Goblin(t)
    g.Describe("测试 compress.go 中的 md5加密 方法", func() {
        var tests = []struct {
            raw string
            md5 string
        }{
            {"j1XngQxn3f", "1d57f2318902faefebbe282d3a816200"},
            {"IjflW2iQDW", "f7802803961e744b02ffb79b9c4036d3"},
            {"信息论、推理与学习算法", "eb7479d1f57cac557550104668366469"},
            {"72400780416952345272", "d882d3ac74b23b668df0f1fa055e296e"},
            {"mqnqjnysyttofubdwikt", "929f7a04d7eef98e081000e2173eb407"},
            {`#@$&@,{()]+#_.>!["]/`, "31d3d310d4e04842b1f9e5a81fc8ef0d"},
            {"HENXEHMAKDPRJIYXTWSD", "a6a8e69ab3cd97c9a32740333f05403e"},
            {"시계열분석", "4b21fb8c79f58e8c498c67ddc7ec3275"},
        }
        g.It("字符串正确地被md5加密", func() {
            for _, test := range tests {
                g.Assert(MD5(test.raw)).Equal(test.md5)
            }
        })
    })

}
