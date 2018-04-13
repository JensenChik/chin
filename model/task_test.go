package model

import (
    "testing"
    . "github.com/franela/goblin"
    "github.com/jinzhu/gorm"
    "time"
    "strconv"
    "../tools/random"
)

func TestTask(t *testing.T) {
    g := Goblin(t)

    g.Describe("测试 task 表插入", func() {
        var db *gorm.DB
        var err error
        var mysqlCount int
        var tasks []Task

        g.Before(func() {
            db, err = connectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            Truncate("tasks")
            tasks = []Task{
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
            }

        })

        g.After(func() {
            defer db.Close()
        })

        g.It(" 记录被正确插入", func() {
            expectedCount := 0
            for id, task := range tasks {
                ok, err := task.DumpToMySQL()
                g.Assert(ok).IsTrue()
                g.Assert(err == nil).IsTrue()

                expectedCount++
                db.Table("tasks").Count(&mysqlCount)
                g.Assert(expectedCount).Equal(mysqlCount)

                db.Model(new(Task)).Where("id = ?", id + 1).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)

                newTask, err := new(Task).LoadByWhere("task_name =?", task.TaskName)
                g.Assert(err == nil)
                g.Assert(newTask.TaskName).Equal(task.TaskName)
                g.Assert(newTask.Command).Equal(task.Command)
                g.Assert(newTask.FatherTask).Equal(task.FatherTask)
                g.Assert(newTask.Valid).Equal(task.Valid)
                g.Assert(newTask.MachinePool).Equal(task.MachinePool)
                g.Assert(newTask.OwnerID).Equal(task.OwnerID)
                g.Assert(newTask.ScheduleType).Equal(task.ScheduleType)
                g.Assert(newTask.ScheduleFormat).Equal(task.ScheduleFormat)

                newTask, err = new(Task).LoadByKey(id + 1)
                g.Assert(err == nil)
                g.Assert(newTask.TaskName).Equal(task.TaskName)
                g.Assert(newTask.Command).Equal(task.Command)
                g.Assert(newTask.FatherTask).Equal(task.FatherTask)
                g.Assert(newTask.Valid).Equal(task.Valid)
                g.Assert(newTask.MachinePool).Equal(task.MachinePool)
                g.Assert(newTask.OwnerID).Equal(task.OwnerID)
                g.Assert(newTask.ScheduleType).Equal(task.ScheduleType)
                g.Assert(newTask.ScheduleFormat).Equal(task.ScheduleFormat)
            }
        })
    })

    g.Describe("测试 task 表更新", func() {
        var db *gorm.DB
        var err error
        var mysqlCount int
        var tasks []Task

        g.Before(func() {
            db, err = connectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            Truncate("tasks")
            tasks = []Task{
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
            }
        })

        g.After(func() {
            defer db.Close()
        })

        g.It("记录被正确更新", func() {
            g.Timeout(10 * time.Second)
            for id, task := range tasks {
                ok, err := task.DumpToMySQL()
                g.Assert(err == nil).IsTrue()
                g.Assert(ok).IsTrue()

                db.Model(new(Task)).Where("task_name = ?", task.TaskName).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)

                oldUpdateTime := task.UpdatedAt
                oldCreateTime := task.CreatedAt

                time.Sleep(time.Second)
                newTaskName := random.String(20)
                newCommand := random.String(20)
                newFatherTask := random.String(20)
                newValid := !task.Valid
                newMachinePool := random.String(20)
                newOwnerID := random.Int(1000)
                newScheduleType := random.String(10)
                newScheduleFormat := random.String(20)
                task.TaskName = newTaskName
                task.Command = newCommand
                task.FatherTask = newFatherTask
                task.Valid = newValid
                task.MachinePool = newMachinePool
                task.OwnerID = newOwnerID
                task.ScheduleType = newScheduleType
                task.ScheduleFormat = newScheduleFormat
                task.DumpToMySQL()

                newTask, err := new(Task).LoadByWhere("id = ?", id + 1)
                g.Assert(err == nil).IsTrue()
                g.Assert(newTask.TaskName).Equal(task.TaskName)
                g.Assert(newTask.Command).Equal(task.Command)
                g.Assert(newTask.FatherTask).Equal(task.FatherTask)
                g.Assert(newTask.Valid).Equal(task.Valid)
                g.Assert(newTask.MachinePool).Equal(task.MachinePool)
                g.Assert(newTask.OwnerID).Equal(task.OwnerID)
                g.Assert(newTask.ScheduleType).Equal(task.ScheduleType)
                g.Assert(newTask.ScheduleFormat).Equal(task.ScheduleFormat)

                g.Assert(newTask.UpdatedAt.Sub(oldUpdateTime).Seconds() > 0).IsTrue()
                g.Assert(newTask.CreatedAt.Format("2006-01-02 15:04:05")).Equal(oldCreateTime.Format("2006-01-02 15:04:05"))
            }

        })
    })

    g.Describe("测试 task 数据加载", func() {
        var db *gorm.DB
        var err error
        var tasks []Task

        g.Before(func() {
            db, err = connectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            Truncate("tasks")
            tasks = []Task{
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
                {TaskName:random.String(20), Command:random.String(10), FatherTask:random.String(10), Valid:random.Int(1) == 0, MachinePool:random.String(10), OwnerID:random.Int(1000), ScheduleType:random.String(10), ScheduleFormat:random.String(10)},
            }
            for _, task := range tasks {
                ok, err := task.DumpToMySQL()
                g.Assert(ok).IsTrue()
                g.Assert(err == nil).IsTrue()
            }
        })

        g.After(func() {
            defer db.Close()
        })

        g.It("记录通过where条件被正确加载", func() {
            for _, task := range tasks {
                newTask, err := new(Task).LoadByWhere("task_name = ?", task.TaskName)
                g.Assert(err == nil).IsTrue()
                g.Assert(newTask.TaskName).Equal(task.TaskName)
                g.Assert(newTask.Command).Equal(task.Command)
                g.Assert(newTask.FatherTask).Equal(task.FatherTask)
                g.Assert(newTask.Valid).Equal(task.Valid)
                g.Assert(newTask.MachinePool).Equal(task.MachinePool)
                g.Assert(newTask.OwnerID).Equal(task.OwnerID)
                g.Assert(newTask.ScheduleType).Equal(task.ScheduleType)
                g.Assert(newTask.ScheduleFormat).Equal(task.ScheduleFormat)
            }
        })

        g.It("记录通主键被正确加载", func() {
            for id, task := range tasks {
                newTask, err := new(Task).LoadByKey(id + 1)
                g.Assert(err == nil).IsTrue()
                g.Assert(newTask.TaskName).Equal(task.TaskName)
                g.Assert(newTask.Command).Equal(task.Command)
                g.Assert(newTask.FatherTask).Equal(task.FatherTask)
                g.Assert(newTask.Valid).Equal(task.Valid)
                g.Assert(newTask.MachinePool).Equal(task.MachinePool)
                g.Assert(newTask.OwnerID).Equal(task.OwnerID)
                g.Assert(newTask.ScheduleType).Equal(task.ScheduleType)
                g.Assert(newTask.ScheduleFormat).Equal(task.ScheduleFormat)
            }
        })

        g.It("记录通过多个where条件被正确加载", func() {
            for id, task := range tasks {
                newTask, err := new(Task).LoadByWhere(
                    "id = ? and task_name = ? and command = ? and father_task = ? and valid = ? and machine_pool = ? and owner_id = ? and schedule_type = ? and schedule_format = ?",
                    id + 1, task.TaskName, task.Command, task.FatherTask, task.Valid, task.MachinePool, task.OwnerID, task.ScheduleType, task.ScheduleFormat,
                )
                g.Assert(err == nil).IsTrue()
                g.Assert(newTask.TaskName).Equal(task.TaskName)
                g.Assert(newTask.Command).Equal(task.Command)
                g.Assert(newTask.FatherTask).Equal(task.FatherTask)
                g.Assert(newTask.Valid).Equal(task.Valid)
                g.Assert(newTask.MachinePool).Equal(task.MachinePool)
                g.Assert(newTask.OwnerID).Equal(task.OwnerID)
                g.Assert(newTask.ScheduleType).Equal(task.ScheduleType)
                g.Assert(newTask.ScheduleFormat).Equal(task.ScheduleFormat)
            }
        })

        g.It("当存在多于一条记录满足where条件时无法实例化，返回异常且对象为nil", func() {
            (&Task{TaskName:"dup_task"}).DumpToMySQL()
            (&Task{TaskName:"dup_task"}).DumpToMySQL()
            task, err := new(Task).LoadByWhere("task_name = ?", "dup_task")
            g.Assert(task == nil).IsTrue()
            g.Assert(err.Error()).Equal("存在多条满足条件的记录，无法实例化")
        })

        g.It("当存在零条记录满足where条件时无法实例化，返回异常且对象为nil", func() {
            newTask, err := new(Task).LoadByWhere("task_name = ?", "大傻逼")
            g.Assert(newTask == nil).IsTrue()
            g.Assert(err.Error()).Equal("不存在满足条件的记录，无法实例化")
        })

    })

    g.Describe("测试 ScheduleToday", func() {
        var task Task
        WEEKDAY_MAPPING := map[string]int{
            "Any": 0,
            "Monday":1,
            "Tuesday":2,
            "Wednesday":3,
            "Thursday":4,
            "Friday":5,
            "Saturday":6,
            "Sunday":7,
        }
        today := time.Now()
        yesterday := today.AddDate(0, 0, -1)
        tomorrow := today.AddDate(0, 0, 1)
        theDayAfterTomorrow := today.AddDate(0, 0, 2)

        g.It("不调度", func() {
            task = Task{Valid:false}
            g.Assert(task.ShouldScheduleToday()).IsFalse()
        })

        g.It("日调度", func() {
            task = Task{ScheduleType:"day", Valid:true}
            g.Assert(task.ShouldScheduleToday()).IsTrue()
            task.Valid = false
            g.Assert(task.ShouldScheduleToday()).IsFalse()
        })

        g.It("周调度", func() {
            todayWeekday := WEEKDAY_MAPPING[today.Weekday().String()]
            yesterdayWeekday := WEEKDAY_MAPPING[yesterday.Weekday().String()]
            tomorrowWeekday := WEEKDAY_MAPPING[tomorrow.Weekday().String()]
            theDayAfterTomorrowWeekday := WEEKDAY_MAPPING[theDayAfterTomorrow.Weekday().String()]
            task = Task{
                Valid:true,
                ScheduleType:"week",
                ScheduleFormat: strconv.Itoa(todayWeekday) + " 0000-00-00 15:04:05",
            }
            g.Assert(task.ShouldScheduleToday()).IsTrue()

            task.Valid = false
            g.Assert(task.ShouldScheduleToday()).IsFalse()

            task.Valid = true

            task.ScheduleType = strconv.Itoa(yesterdayWeekday) + " 0000-00-00 15:04:05"
            g.Assert(task.ShouldScheduleToday()).IsFalse()

            task.ScheduleType = strconv.Itoa(tomorrowWeekday) + " 0000-00-00 15:04:05"
            g.Assert(task.ShouldScheduleToday()).IsFalse()

            task.ScheduleType = strconv.Itoa(theDayAfterTomorrowWeekday) + " 0000-00-00 15:04:05"
            g.Assert(task.ShouldScheduleToday()).IsFalse()
        })

        g.It("月调度", func() {
            task = Task{
                Valid:true,
                ScheduleType:"month",
                ScheduleFormat:"0 0000-00-" + today.Format("02") + " 15:04:05",
            }
            g.Assert(task.ShouldScheduleToday()).IsTrue()

            task.Valid = false
            g.Assert(task.ShouldScheduleToday()).IsFalse()

            task.Valid = true

            task.ScheduleFormat = "0 0000-00-" + yesterday.Format("02") + " 15:04:05"
            g.Assert(task.ShouldScheduleToday()).IsFalse()

            task.ScheduleFormat = "0 0000-00-" + tomorrow.Format("02") + " 15:04:05"
            g.Assert(task.ShouldScheduleToday()).IsFalse()

            task.ScheduleFormat = "0 0000-00-" + theDayAfterTomorrow.Format("02") + " 15:04:05"
            g.Assert(task.ShouldScheduleToday()).IsFalse()

        })

        g.It("单次调度", func() {
            task = Task{
                Valid:true,
                ScheduleType:"month",
                ScheduleFormat:"0 " + today.Format("2006-01-02") + " 15:04:05",
            }
            g.Assert(task.ShouldScheduleToday()).IsTrue()

            task.Valid = false
            g.Assert(task.ShouldScheduleToday()).IsFalse()

            task.Valid = true

            task.ScheduleFormat = "0 " + yesterday.Format("2006-01-02") + " 15:04:05"
            g.Assert(task.ShouldScheduleToday()).IsFalse()

            task.ScheduleFormat = "0 " + tomorrow.Format("2006-01-02") + " 15:04:05"
            g.Assert(task.ShouldScheduleToday()).IsFalse()

            task.ScheduleFormat = "0 " + theDayAfterTomorrow.Format("2006-01-02") + " 15:04:05"
            g.Assert(task.ShouldScheduleToday()).IsFalse()
        })

    })

    g.Describe("测试 NoJobToday", func() {
        var db *gorm.DB
        var err error
        var task *Task
        g.BeforeEach(func() {
            db, err = connectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            Truncate("jobs")
            Truncate("tasks")
        })

        g.After(func() {
            defer db.Close()
        })

        g.It("正确地判断是否 NoJobToday", func() {
            taskID := random.Int(100)
            for i := 0; i < int(taskID); i++ {
                task = new(Task)
                task.DumpToMySQL()
            }
            job := Job{TaskID:taskID}
            job.CreatedAt = time.Now().AddDate(0, 0, -1)
            job.DumpToMySQL()
            g.Assert(task.NoJobToday()).IsTrue()
            job = Job{TaskID:taskID}
            job.DumpToMySQL()
            g.Assert(task.NoJobToday()).IsFalse()
        })

    })

    g.Describe("测试 SuccessToday", func() {
        g.It("若当天不应调度则直接返回false", func() {
            task := Task{
                Valid: false,
                ScheduleType:"once",
                ScheduleFormat:"0 " + time.Now().AddDate(0, 0, -1).Format("2006-01-02") + " 00:10:00",
            }
            g.Assert(task.ShouldScheduleToday()).IsFalse()
            g.Assert(task.SuccessToday()).IsFalse()
        })

        g.It("若当天对应的 job 状态不为 success 则返回false", func() {
            task := Task{
                Valid:true,
                ScheduleType:"once",
                ScheduleFormat:"0 " + time.Now().Format("2006-01-02") + " 00:10:00",
            }
            g.Assert(task.ShouldScheduleToday()).IsTrue()
            g.Assert(task.NoJobToday()).IsTrue()
            task.CreateJob()
            g.Assert(task.NoJobToday()).IsFalse()
            job := new(Job)
            job.LoadByWhere("task_id = ?", task.ID)

            job.Status = "pooling"
            job.DumpToMySQL()
            g.Assert(task.SuccessToday()).IsFalse()

            job.Status = "failed"
            job.DumpToMySQL()
            g.Assert(task.SuccessToday()).IsFalse()

            job.Status = "success"
            job.DumpToMySQL()
            g.Assert(task.SuccessToday()).IsTrue()



        })

    })

}