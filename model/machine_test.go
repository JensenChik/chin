package model

import (
    "testing"
    . "github.com/franela/goblin"
    "github.com/jinzhu/gorm"
    "time"
    "math"
    "../tools/random"
)

func TestMachine(t *testing.T) {
    g := Goblin(t)

    g.Describe("测试 machine 表插入", func() {
        var db *gorm.DB
        var err error
        var mysqlCount int
        var machines []Machine

        g.Before(func() {
            db, err = connectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            Truncate("machines")
            machines = []Machine{
                {CPULoad:random.Float(), MemoryLoad:random.Float(), MachineName:random.String(20), IP:random.String(12), MAC:random.String(10), Alive:random.Int(1) == 1},
                {CPULoad:random.Float(), MemoryLoad:random.Float(), MachineName:random.String(20), IP:random.String(12), MAC:random.String(10), Alive:random.Int(1) == 1},
                {CPULoad:random.Float(), MemoryLoad:random.Float(), MachineName:random.String(20), IP:random.String(12), MAC:random.String(10), Alive:random.Int(1) == 1},
                {CPULoad:random.Float(), MemoryLoad:random.Float(), MachineName:random.String(20), IP:random.String(12), MAC:random.String(10), Alive:random.Int(1) == 1},
                {CPULoad:random.Float(), MemoryLoad:random.Float(), MachineName:random.String(20), IP:random.String(12), MAC:random.String(10), Alive:random.Int(1) == 1},
                {CPULoad:random.Float(), MemoryLoad:random.Float(), MachineName:random.String(20), IP:random.String(12), MAC:random.String(10), Alive:random.Int(1) == 1},
            }

        })

        g.After(func() {
            defer db.Close()
        })

        g.It(" 记录被正确插入", func() {
            expectedCount := 0
            for id, machine := range machines {
                ok, err := machine.DumpToMySQL()
                g.Assert(ok).IsTrue()
                g.Assert(err == nil).IsTrue()

                expectedCount++
                db.Table("machines").Count(&mysqlCount)
                g.Assert(expectedCount).Equal(mysqlCount)

                db.Model(new(Machine)).Where("id = ?", id + 1).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)

                newMachine, err := new(Machine).LoadByWhere("machine_name =?", machine.MachineName)
                g.Assert(err == nil)
                g.Assert(newMachine.IP).Equal(machine.IP)
                g.Assert(newMachine.MAC).Equal(machine.MAC)
                g.Assert(math.Abs(newMachine.CPULoad - machine.CPULoad) < 1e-6).IsTrue()
                g.Assert(math.Abs(newMachine.MemoryLoad - machine.MemoryLoad) < 1e-6).IsTrue()
                g.Assert(newMachine.Alive).Equal(machine.Alive)

                newMachine, err = new(Machine).LoadByKey(id + 1)
                g.Assert(err == nil)
                g.Assert(newMachine.MachineName).Equal(machine.MachineName)
                g.Assert(newMachine.IP).Equal(machine.IP)
                g.Assert(newMachine.MAC).Equal(machine.MAC)
                g.Assert(newMachine.CPULoad).Equal(machine.CPULoad)
                g.Assert(newMachine.MemoryLoad).Equal(machine.MemoryLoad)
                g.Assert(newMachine.Alive).Equal(machine.Alive)
            }
        })
    })

    g.Describe("测试 machine 表更新", func() {
        var db *gorm.DB
        var err error
        var mysqlCount int
        var machines []Machine

        g.Before(func() {
            db, err = connectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            Truncate("machines")
            machines = []Machine{
                {CPULoad:random.Float(), MemoryLoad:random.Float(), MachineName:random.String(20), IP:random.String(12), MAC:random.String(10), Alive:random.Int(1) == 1},
                {CPULoad:random.Float(), MemoryLoad:random.Float(), MachineName:random.String(20), IP:random.String(12), MAC:random.String(10), Alive:random.Int(1) == 1},
                {CPULoad:random.Float(), MemoryLoad:random.Float(), MachineName:random.String(20), IP:random.String(12), MAC:random.String(10), Alive:random.Int(1) == 1},
                {CPULoad:random.Float(), MemoryLoad:random.Float(), MachineName:random.String(20), IP:random.String(12), MAC:random.String(10), Alive:random.Int(1) == 1},
                {CPULoad:random.Float(), MemoryLoad:random.Float(), MachineName:random.String(20), IP:random.String(12), MAC:random.String(10), Alive:random.Int(1) == 1},
                {CPULoad:random.Float(), MemoryLoad:random.Float(), MachineName:random.String(20), IP:random.String(12), MAC:random.String(10), Alive:random.Int(1) == 1},
            }
        })

        g.After(func() {
            defer db.Close()
        })

        g.It("记录被正确更新", func() {
            g.Timeout(10 * time.Second)
            for id, machine := range machines {
                ok, err := machine.DumpToMySQL()
                g.Assert(err == nil).IsTrue()
                g.Assert(ok).IsTrue()

                db.Model(new(Machine)).Where("machine_name = ?", machine.MachineName).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)

                oldUpdateTime := machine.UpdatedAt
                oldCreateTime := machine.CreatedAt

                time.Sleep(time.Second)
                newMachineName := random.String(10)
                newIP := random.String(10)
                newMAC := random.String(17)
                newCPULoad := random.Float()
                newMemoryLoad := random.Float()
                newAlive := !machine.Alive

                machine.MachineName = newMachineName
                machine.IP = newIP
                machine.MAC = newMAC
                machine.CPULoad = newCPULoad
                machine.MemoryLoad = newMemoryLoad
                machine.Alive = newAlive

                machine.DumpToMySQL()

                newMachine, _ := new(Machine).LoadByWhere("id = ?", id + 1)
                g.Assert(newMachine.MachineName).Equal(newMachineName)
                g.Assert(newMachine.IP).Equal(newIP)
                g.Assert(newMachine.MAC).Equal(newMAC)
                g.Assert(math.Abs(newMachine.CPULoad - machine.CPULoad) < 1e-6).IsTrue()
                g.Assert(math.Abs(newMachine.MemoryLoad - machine.MemoryLoad) < 1e-6).IsTrue()
                g.Assert(newMachine.Alive).Equal(newAlive)
                g.Assert(newMachine.UpdatedAt.Sub(oldUpdateTime).Seconds() > 0).IsTrue()
                g.Assert(newMachine.CreatedAt.Format("2006-01-02 15:04:05")).Equal(oldCreateTime.Format("2006-01-02 15:04:05"))
            }

        })
    })

    g.Describe("测试 machine 数据加载", func() {
        var db *gorm.DB
        var err error
        var machines []Machine

        g.Before(func() {
            db, err = connectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            Truncate("machines")
            machines = []Machine{
                {CPULoad:random.Float(), MemoryLoad:random.Float(), MachineName:random.String(20), IP:random.String(12), MAC:random.String(10), Alive:random.Int(1) == 1},
                {CPULoad:random.Float(), MemoryLoad:random.Float(), MachineName:random.String(20), IP:random.String(12), MAC:random.String(10), Alive:random.Int(1) == 1},
                {CPULoad:random.Float(), MemoryLoad:random.Float(), MachineName:random.String(20), IP:random.String(12), MAC:random.String(10), Alive:random.Int(1) == 1},
                {CPULoad:random.Float(), MemoryLoad:random.Float(), MachineName:random.String(20), IP:random.String(12), MAC:random.String(10), Alive:random.Int(1) == 1},
                {CPULoad:random.Float(), MemoryLoad:random.Float(), MachineName:random.String(20), IP:random.String(12), MAC:random.String(10), Alive:random.Int(1) == 1},
                {CPULoad:random.Float(), MemoryLoad:random.Float(), MachineName:random.String(20), IP:random.String(12), MAC:random.String(10), Alive:random.Int(1) == 1},
            }

            for _, machine := range machines {
                ok, err := machine.DumpToMySQL()
                g.Assert(ok).IsTrue()
                g.Assert(err == nil).IsTrue()
            }
        })

        g.After(func() {
            defer db.Close()
        })

        g.It("记录通过where条件被正确加载", func() {
            for _, machine := range machines {
                newMachine, err := new(Machine).LoadByWhere("machine_name = ?", machine.MachineName)
                g.Assert(err == nil).IsTrue()
                g.Assert(newMachine.MachineName).Equal(machine.MachineName)
                g.Assert(newMachine.IP).Equal(machine.IP)
                g.Assert(newMachine.MAC).Equal(machine.MAC)
                g.Assert(math.Abs(newMachine.CPULoad - machine.CPULoad) < 1e-6).IsTrue()
                g.Assert(math.Abs(newMachine.MemoryLoad - machine.MemoryLoad) < 1e-6).IsTrue()
                g.Assert(newMachine.Alive).Equal(machine.Alive)
            }
        })

        g.It("记录通主键被正确加载", func() {
            for id, machine := range machines {
                newMachine, err := new(Machine).LoadByKey(id + 1)
                g.Assert(err == nil).IsTrue()
                g.Assert(newMachine.MachineName).Equal(machine.MachineName)
                g.Assert(newMachine.IP).Equal(machine.IP)
                g.Assert(newMachine.MAC).Equal(machine.MAC)
                g.Assert(math.Abs(newMachine.CPULoad - machine.CPULoad) < 1e-6).IsTrue()
                g.Assert(math.Abs(newMachine.MemoryLoad - machine.MemoryLoad) < 1e-6).IsTrue()
                g.Assert(newMachine.Alive).Equal(machine.Alive)
            }
        })

        g.It("记录通过多个where条件被正确加载", func() {
            for id, machine := range machines {
                newMachine, err := new(Machine).LoadByWhere(
                    "id = ? and machine_name = ? and ip = ? and mac = ? and alive = ?",
                    id + 1, machine.MachineName, machine.IP, machine.MAC,  machine.Alive,
                )
                g.Assert(err == nil).IsTrue()
                g.Assert(newMachine.MachineName).Equal(machine.MachineName)
                g.Assert(newMachine.IP).Equal(machine.IP)
                g.Assert(newMachine.MAC).Equal(machine.MAC)
                g.Assert(newMachine.CPULoad).Equal(machine.CPULoad)
                g.Assert(newMachine.MemoryLoad).Equal(machine.MemoryLoad)
                g.Assert(newMachine.Alive).Equal(machine.Alive)
            }
        })

        g.It("当存在多于一条记录满足where条件时无法实例化，返回异常且对象为nil", func() {
            (&Machine{MachineName:"dup_machine"}).DumpToMySQL()
            (&Machine{MachineName:"dup_machine"}).DumpToMySQL()
            machine, err := new(Machine).LoadByWhere("machine_name = ?", "dup_machine")
            g.Assert(machine == nil).IsTrue()
            g.Assert(err.Error()).Equal("存在多条满足条件的记录，无法实例化")
        })

        g.It("当存在零条记录满足where条件时无法实例化，返回异常且对象为nil", func() {
            newMachine, err := new(Machine).LoadByWhere("machine_name = ?", "大傻逼")
            g.Assert(newMachine == nil).IsTrue()
            g.Assert(err.Error()).Equal("不存在满足条件的记录，无法实例化")
        })

    })
}