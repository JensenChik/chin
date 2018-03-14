package worker

import (
    "testing"
    . "github.com/franela/goblin"
)

func TestSys(t *testing.T) {
    g := Goblin(t)

    g.Describe("sys 单元测试", func() {
        var stat *sysStat
        g.BeforeEach(func() {
            stat = new(sysStat)
        })

        g.It("机器负载", func() {
            stat.loadAvg()
            g.Assert(stat.Load1 > 0 && stat.Load1 < 1).IsTrue()
            g.Assert(stat.Load5 > 0 && stat.Load5 < 1).IsTrue()
            g.Assert(stat.Load15 > 0 && stat.Load15 < 1).IsTrue()
        })

        g.It("虚拟内存", func() {
            stat.virtualMemory()
            g.Assert(stat.MemActive > 0).IsTrue()
            g.Assert(stat.MemAvailable > 0).IsTrue()
            g.Assert(stat.MemBuffer > 0).IsTrue()
            g.Assert(stat.MemCache > 0).IsTrue()
            g.Assert(stat.MemFree > 0).IsTrue()
            g.Assert(stat.MemInactive > 0).IsTrue()
            g.Assert(stat.MemTotal > 0).IsTrue()
            g.Assert(stat.MemUsed > 0).IsTrue()
            g.Assert(stat.MemUsedPercent > 0 && stat.MemUsedPercent < 100).IsTrue()
        })

        g.It("网络流量", func() {
            stat.network()
            g.Assert(stat.NetSendByte > 0).IsTrue()
            g.Assert(stat.NetSendPack > 0).IsTrue()
            g.Assert(stat.NetRecvByte > 0).IsTrue()
            g.Assert(stat.NetRecvPack > 0).IsTrue()
        })


    })

}
