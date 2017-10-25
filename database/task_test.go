package database

import (
    "testing"
    . "github.com/franela/goblin"
    "github.com/jinzhu/gorm"
    "time"
)

func TestTask(t *testing.T) {
    g := Goblin(t)

    g.Describe("测试 task 表插入", func() {
        var db *gorm.DB
        var err error
        var mysqlCount int
        var tasks []Task

        g.Before(func() {
            db, err = ConnectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            db.Exec("truncate table tasks;")
            tasks = []Task{
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
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
            db, err = ConnectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            db.Exec("truncate table tasks;")
            tasks = []Task{
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
            }
        })

        g.After(func() {
            defer db.Close()
        })

        g.Xit("记录被正确更新", func() {
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
                newTaskName := randomString(20)
                newCommand := randomString(20)
                newFatherTask := randomString(20)
                newValid := !task.Valid
                newMachinePool := randomString(20)
                newOwnerID := randomInt(1000)
                newScheduleType := randomString(10)
                newScheduleFormat := randomString(20)
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
            db, err = ConnectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            db.Exec("truncate table tasks;")
            tasks = []Task{
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
                {TaskName:randomString(20), Command:randomString(10), FatherTask:randomString(10), Valid:randomInt(1) == 0, MachinePool:randomString(10), OwnerID:randomInt(1000), ScheduleType:randomString(10), ScheduleFormat:randomString(10)},
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
}