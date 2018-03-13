package worker

import (
    "testing"
    . "github.com/franela/goblin"
)

func TestSys(t *testing.T) {
    g := Goblin(t)

    g.Describe("loadAvg 单元测试", func() {
        g.It("load1, load5, load15应为浮点数", func() {
            stat := new(sysStat)
            stat.loadAvg()
            g.Assert(stat.Load1 > 0 && stat.Load1 < 1).IsTrue()
            g.Assert(stat.Load5 > 0 && stat.Load5 < 1).IsTrue()
            g.Assert(stat.Load15 > 0 && stat.Load15 < 1).IsTrue()
        })
    })
}
