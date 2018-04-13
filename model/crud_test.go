package model

import (
    "testing"
    . "github.com/franela/goblin"
    "github.com/jinzhu/gorm"
    "../tools/random"
)

func TestCRUD(t *testing.T) {
    g := Goblin(t)
    g.Describe("select all where", func() {
        var db *gorm.DB
        var err error
        var jobs []Job

        g.Before(func() {
            db, err = connectDatabase()
            if err != nil {
                g.Fail("连接mysql错误")
            }
            db.Exec("truncate table jobs;")
            jobs = []Job{
                {TaskID:100, Status:random.String(20)},
                {TaskID:200, Status:random.String(20)},
                {TaskID:300, Status:random.String(20)},
                {TaskID:400, Status:random.String(20)},
                {TaskID:500, Status:random.String(20)},
                {TaskID:600, Status:random.String(20)},
                {TaskID:700, Status:random.String(20)},
                {TaskID:800, Status:random.String(20)},
                {TaskID:900, Status:random.String(20)},
                {TaskID:1000, Status:random.String(20)},
            }
            for _, job := range jobs {
                job.DumpToMySQL()
            }
        })

        g.It("Truncate测试")

        g.It("通过单个where条件过滤", func() {
            Fill(&jobs).Where("id <= ?", 10)
            g.Assert(len(jobs)).Equal(10)
            Fill(&jobs).Where("id <= ?", 0)
            g.Assert(len(jobs)).Equal(0)
            Fill(&jobs).Where("id > ?", 10)
            g.Assert(len(jobs)).Equal(0)
        })

        g.It("通过多个where条件过滤", func() {
            Fill(&jobs).Where("id <= ? and task_id <= ?", 10, 300)
            g.Assert(len(jobs)).Equal(3)
            Fill(&jobs).Where("id <= ? and task_id <= ?", 0, 300)
            g.Assert(len(jobs)).Equal(0)
            Fill(&jobs).Where("id <= ? and task_id <= ?", 10, 0)
            g.Assert(len(jobs)).Equal(0)
            Fill(&jobs).Where("id <= ? and task_id <= ?", 0, 0)
            g.Assert(len(jobs)).Equal(0)
        })
    })
}
