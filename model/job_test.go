package model

import (
    "testing"
    "time"

    "../tools/random"
    . "github.com/franela/goblin"
    "github.com/jinzhu/gorm"
)

func TestJob(t *testing.T) {
    g := Goblin(t)

    g.Describe("测试 jobs 表插入", func() {
        var db *gorm.DB
        var err error
        var mysqlCount int
        var jobs []Job

        g.Before(func() {
            db, err = connectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            Truncate("jobs")
            jobs = []Job{
                {TaskID: random.Int(10000), Status: random.String(20)},
                {TaskID: random.Int(10000), Status: random.String(20)},
                {TaskID: random.Int(10000), Status: random.String(20)},
                {TaskID: random.Int(10000), Status: random.String(20)},
                {TaskID: random.Int(10000), Status: random.String(20)},
                {TaskID: random.Int(10000), Status: random.String(20)},
            }

        })

        g.After(func() {
            defer db.Close()
        })

        g.It("jobs记录被正确插入", func() {
            expectedCount := 0
            for id, job := range jobs {
                ok, err := job.DumpToMySQL()
                g.Assert(ok).IsTrue()
                g.Assert(err == nil).IsTrue()

                expectedCount++
                db.Table("jobs").Count(&mysqlCount)
                g.Assert(expectedCount).Equal(mysqlCount)

                db.Model(new(Job)).Where("id = ?", id+1).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)

                job, err := new(Job).LoadByWhere("task_id =?", job.TaskID)
                g.Assert(err == nil)
                g.Assert(job.Status).Equal(job.Status)

                job, err = new(Job).LoadByKey(id + 1)
                g.Assert(err == nil)
                g.Assert(job.TaskID).Equal(job.TaskID)
                g.Assert(job.Status).Equal(job.Status)
            }
        })
    })

    g.Describe("测试 jobs 表更新", func() {
        var db *gorm.DB
        var err error
        var mysqlCount int
        var jobs []Job

        g.Before(func() {
            db, err = connectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            Truncate("jobs")
            jobs = []Job{
                {TaskID: random.Int(10000), Status: random.String(20)},
                {TaskID: random.Int(10000), Status: random.String(20)},
                {TaskID: random.Int(10000), Status: random.String(20)},
                {TaskID: random.Int(10000), Status: random.String(20)},
                {TaskID: random.Int(10000), Status: random.String(20)},
                {TaskID: random.Int(10000), Status: random.String(20)},
            }
        })

        g.After(func() {
            defer db.Close()
        })

        g.It("记录被正确更新", func() {
            g.Timeout(10 * time.Second)
            for _, job := range jobs {
                ok, err := job.DumpToMySQL()
                g.Assert(err == nil).IsTrue()
                g.Assert(ok).IsTrue()

                db.Model(new(Job)).Where("task_id = ?", job.TaskID).Count(&mysqlCount)
                g.Assert(mysqlCount).Equal(1)

                oldUpdateTime := job.UpdatedAt
                oldCreateTime := job.CreatedAt

                time.Sleep(time.Second)
                newStatus := random.String(3)
                job.Status = newStatus
                job.DumpToMySQL()

                newJob, _ := new(Job).LoadByWhere("task_id = ?", job.TaskID)
                g.Assert(newJob.Status).Equal(newStatus)
                g.Assert(newJob.UpdatedAt.Sub(oldUpdateTime).Seconds() > 0).IsTrue()
                g.Assert(newJob.CreatedAt.Format("2006-01-02 15:04:05")).Equal(oldCreateTime.Format("2006-01-02 15:04:05"))
            }

        })
    })

    g.Describe("测试 jobs 数据加载", func() {
        var db *gorm.DB
        var err error
        var jobs []Job

        g.Before(func() {
            db, err = connectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            Truncate("jobs")
            jobs = []Job{
                {TaskID: random.Int(10000), Status: random.String(20)},
                {TaskID: random.Int(10000), Status: random.String(20)},
                {TaskID: random.Int(10000), Status: random.String(20)},
                {TaskID: random.Int(10000), Status: random.String(20)},
                {TaskID: random.Int(10000), Status: random.String(20)},
                {TaskID: random.Int(10000), Status: random.String(20)},
            }

            for _, job := range jobs {
                ok, err := job.DumpToMySQL()
                g.Assert(ok).IsTrue()
                g.Assert(err == nil).IsTrue()
            }
        })

        g.After(func() {
            defer db.Close()
        })

        g.It("记录通过where条件被正确加载", func() {
            for _, job := range jobs {
                jewJob, err := new(Job).LoadByWhere("task_id = ?", job.TaskID)
                g.Assert(err == nil).IsTrue()
                g.Assert(jewJob.Status).Equal(job.Status)
            }
        })

        g.It("记录通主键被正确加载", func() {
            for id, job := range jobs {
                newJob, err := new(Job).LoadByKey(id + 1)
                g.Assert(err == nil).IsTrue()
                g.Assert(newJob.TaskID).Equal(job.TaskID)
                g.Assert(newJob.Status).Equal(job.Status)
            }
        })

        g.It("记录通过多个where条件被正确加载", func() {
            for id, job := range jobs {
                newJob, err := new(Job).LoadByWhere(
                    "id = ? and task_id = ?",
                    id+1, job.TaskID,
                )
                g.Assert(err == nil).IsTrue()
                g.Assert(newJob.Status).Equal(job.Status)
            }
        })

        g.It("当存在多于一条记录满足where条件时无法实例化，返回异常且对象为nil", func() {
            (&Job{TaskID: 12580, Status: random.String(20)}).DumpToMySQL()
            (&Job{TaskID: 12580, Status: random.String(20)}).DumpToMySQL()
            job, err := new(Job).LoadByWhere("task_id = ?", 12580)
            g.Assert(job == nil).IsTrue()
            g.Assert(err.Error()).Equal("存在多条满足条件的记录，无法实例化")
        })

        g.It("当存在零条记录满足where条件时无法实例化，返回异常且对象为nil", func() {
            newJob, err := new(Job).LoadByWhere("task_id = ?", 999999)
            g.Assert(newJob == nil).IsTrue()
            g.Assert(err.Error()).Equal("不存在满足条件的记录，无法实例化")
        })

    })

    g.Describe("测试 job 是否 ready", func() {

    })
}
