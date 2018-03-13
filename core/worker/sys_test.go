package worker

import (
    "testing"
    . "github.com/franela/goblin"
    "github.com/sdbaiguanghe/glog"
)

func TestSys(t *testing.T) {
    g := Goblin(t)

    g.Describe("sys 单元测试", func() {
        var stat *sysStat
        g.BeforeEach(func() {
            stat = new(sysStat)
        })

        g.It("load1, load5, load15应为浮点数", func() {
            stat.loadAvg()
            g.Assert(stat.Load1 > 0 && stat.Load1 < 1).IsTrue()
            g.Assert(stat.Load5 > 0 && stat.Load5 < 1).IsTrue()
            g.Assert(stat.Load15 > 0 && stat.Load15 < 1).IsTrue()
        })

        g.It("虚拟内存", func() {
            stat.virtualMemory()
            glog.Error(stat.MemTotal)

        })
    })

}
