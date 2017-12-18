package database

import (
    "github.com/jinzhu/gorm"
    "strings"
    "time"
)

type Task struct {
    gorm.Model
    TaskName       string
    Command        string
    FatherTask     string `gorm:"type:text"`
    NotifyList     string `gorm:"type:text"`
    Valid          bool
    MachinePool    string
    OwnerID        uint
    RetryTimes     uint
    ScheduleType   string
    ScheduleFormat string // <%u %Y-%m-%d %H:%M:%S>
}

func (task *Task) ShouldScheduleToday() (bool) {
    if !task.Valid {
        return false
    }
    WEEKDAY_MAPPING := map[string]string{
        "Any": "0",
        "Monday":"1",
        "Tuesday":"2",
        "Wednesday":"3",
        "Thursday":"4",
        "Friday":"5",
        "Saturday":"6",
        "Sunday":"7",
    }
    WEEKDAY_IDX := 0
    DATE_IDX := 1
    DAY_IDX := 2
    SPACE := " "
    DASH := "-"
    scheduleToday := false

    switch task.ScheduleType {
    case "day": // 0 0000-00-00 15:04:05
        scheduleToday = true
    case "week": // 1 0000-00-00 15:04:05
        scheduleToday = WEEKDAY_MAPPING[time.Now().Weekday().String()] ==
            strings.Split(task.ScheduleFormat, SPACE)[WEEKDAY_IDX]
    case "month": // 0 0000-00-02 15:04:05
        scheduleToday = time.Now().Format("02") == strings.Split(
            strings.Split(task.ScheduleFormat, SPACE)[DATE_IDX],
            DASH,
        )[DAY_IDX]
    case "once": // 0 2006-01-02 15:04:05
        scheduleToday = time.Now().Format("2006-01-02") ==
            strings.Split(task.ScheduleFormat, SPACE)[DATE_IDX]
    }
    return scheduleToday
}

func (task *Task) NoJobToday() (bool) {
    today := time.Now().Format("2006-01-02")
    jobs := []Job{}
    Fill(&jobs).Where("task_id = ? and date(created_at) = ? ", task.ID, today)
    return len(jobs) == 0
}

func (task *Task) CreateJob() {
    job := new(Job)
    job.TaskID = task.ID
    job.Status = "pooling"
    job.DumpToMySQL()
}

func (task *Task) SuccessToday() bool {
    if !task.ShouldScheduleToday() || task.NoJobToday() {
        return false
    }
    today := time.Now().Format("2006-01-02")
    job := new(Job)
    job.LoadByWhere("task_id = ? and date(created_at) = ? ", task.ID, today)
    return job.Status == "success"
}

func (task *Task) DumpToMySQL() (bool, error) {
    ok, err := dumpToMysql(task)
    return ok, err
}

func (task *Task) LoadByWhere(filters ...interface{}) (*Task, error) {
    initTask, err := loadByWhere(task, filters...)
    if err != nil {
        return nil, err
    } else {
        return initTask.(*Task), nil
    }
}

func (task *Task) LoadByKey(key interface{}) (*Task, error) {
    initTask, err := loadByKey(task, key)
    if err != nil {
        return nil, err
    } else {
        return initTask.(*Task), nil
    }
}
