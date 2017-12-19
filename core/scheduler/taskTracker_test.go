package scheduler

import (
    "testing"
    "../../database"
    . "github.com/franela/goblin"
    "time"
)

func TestTaskTracker(t *testing.T) {
    g := Goblin(t)

    g.Describe("task tracker 单元测试", func() {

        g.BeforeEach(func() {
            database.Truncate("tasks")
            database.Truncate("jobs")
        })

        g.It("为task新建一个job", func() {
            task := database.Task{}
            task.DumpToMySQL()

            jobs := []database.Job{}
            database.Fill(&jobs).Where("task_id = ?", task.ID)
            g.Assert(len(jobs)).Equal(0)

            task.CreateJob()

            database.Fill(&jobs).Where("task_id = ?", task.ID)
            g.Assert(len(jobs)).Equal(1)

        })

        g.It("模拟TaskTracker启动时的跨天", func() {
            var date string
            current := currentDate()
            g.Assert(current == date).IsFalse()
            g.Assert(currentDate).Equal(time.Now().Format("2006-01-02"))
            date = current
            current = currentDate()
            g.Assert(current == date).IsTrue()
            g.Assert(current == "").IsTrue()
        })

        g.It("模拟TaskTracker例行调度的跨天", func() {
            var date = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
            current := currentDate()
            g.Assert(current == date).IsFalse()

            date = current

            current = currentDate()
            g.Assert(current == date).IsTrue()

            date = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
            current = currentDate()
            g.Assert(date == current).IsFalse()

        })

        g.It("模拟taskTracker启动", func() {
            g.Timeout(10 * time.Second)
            var task database.Task
            task = database.Task{Valid:false}
            task.DumpToMySQL()

            task = database.Task{
                Valid:true,
                ScheduleType:"once",
                ScheduleFormat:"0 " + time.Now().Format("2006-01-02") + " 11:00:00",
            }
            task.DumpToMySQL()

            task = database.Task{
                Valid:true,
                ScheduleType:"once",
                ScheduleFormat:"0 " + time.Now().AddDate(0, 0, -1).Format("2006-01-02") + " 11:00:00",
            }
            task.DumpToMySQL()

            task = database.Task{
                Valid:true,
                ScheduleType: "day",
                ScheduleFormat:"0 0000-00-00 11:00:00",
            }
            task.DumpToMySQL()

            go taskTracker()
            time.Sleep(2 * time.Second)
            jobs := []database.Job{}
            database.Fill(&jobs).Where("true")

            g.Assert(len(jobs)).Equal(2)
            g.Assert(jobs[0].TaskID).Equal(uint(2))
            g.Assert(jobs[1].TaskID).Equal(uint(4))

        })

    })
}

