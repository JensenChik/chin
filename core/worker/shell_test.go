package worker

import (
    "testing"
    "time"

    . "github.com/franela/goblin"
)

func TestShell(t *testing.T) {
    g := Goblin(t)
    g.Describe("测试shell", func() {
        g.It("任务尚在运行", func() {
            g.Timeout(10 * time.Second)
            sh := Shell{Command: "sleep 2; echo shell is running"}
            sh.Run()
            time.Sleep(time.Second)
            g.Assert(sh.Finish).IsFalse()
        })

        g.It("任务运行完毕", func() {
            g.Timeout(10 * time.Second)
            sh := Shell{Command: "sleep 1; echo shell finish"}
            sh.Run()
            time.Sleep(2 * time.Second)
            g.Assert(sh.Finish).IsTrue()
        })

        g.It("任务运行成功", func() {
            sh := Shell{Command: "echo shell run success"}
            sh.Run()
            time.Sleep(1 * time.Second)
            g.Assert(sh.Finish).IsTrue()
            g.Assert(sh.Success).IsTrue()
        })

        g.It("任务运行失败", func() {
            sh := Shell{Command: "anyFuckingCommandThatWillNotRun args"}
            sh.Run()
            time.Sleep(1 * time.Second)
            g.Assert(sh.Finish).IsTrue()
            g.Assert(sh.Success).IsFalse()
        })

        g.It("运行成功返回正确的输出", func() {
            sh := Shell{Command: "echo helloworld"}
            sh.Run()
            time.Sleep(1 * time.Second)
            g.Assert(sh.Finish).IsTrue()
            g.Assert(sh.Success).IsTrue()
            g.Assert(sh.Output).Equal("helloworld\n")
        })

        g.It("运行错误返回错误日志", func() {
            sh := Shell{Command: "anyFuckingCommandThatWillNotRun args"}
            sh.Run()
            time.Sleep(1 * time.Second)
            g.Assert(sh.Finish).IsTrue()
            g.Assert(sh.Success).IsFalse()
            g.Assert(sh.Output).Equal("sh: 1: anyFuckingCommandThatWillNotRun: not found\n")
        })

        g.It("任务被杀死", func() {
            sh := Shell{Command: "echo kill; sleep 10"}
            sh.Run()
            time.Sleep(1 * time.Second)
            g.Assert(sh.Finish).IsFalse()
            sh.Kill()
            time.Sleep(1 * time.Second)
            g.Assert(sh.Finish).IsTrue()
            g.Assert(sh.Success).IsFalse()
            g.Assert(sh.Output).Equal("kill\n")
        })
    })

}
