package model

import (
    "github.com/jinzhu/gorm"
    "time"
    "fmt"
    "strings"
    "../tools/datetime"
    "../tools/number"
    "github.com/tidwall/gjson"
    "strconv"
)

type scheduleFormat struct {
    period  string
    weekday int
    year    int
    month   int
    day     int
    hour    int
    minute  int
    second  int
}

type Task struct {
    gorm.Model
    TaskName    string
    Command     string
    FatherTask  string `gorm:"type:text"`
    NotifyList  string `gorm:"type:text"`
    Valid       bool
    MachinePool string
    OwnerID     uint
    RetryTimes  uint
    Schedule    string
    schedule    scheduleFormat `gorm:"-"`
}

func (task *Task) ShouldScheduleToday() bool {
    if !task.Valid {
        return false
    }
    scheduleToday := false
    now := time.Now()
    switch task.schedule.period {
    case "day":
        scheduleToday = true
    case "week":
        scheduleToday = task.schedule.weekday == int(now.Weekday())
    case "month":
        scheduleToday = task.schedule.day == now.Day()
    case "once":
        scheduleToday = task.schedule.year == now.Year() && task.schedule.month == int(now.Month())
    }
    return scheduleToday
}

func (task *Task) ReachScheduleClock() bool {
    now := time.Now()
    return now.Hour() > task.schedule.hour &&
        now.Minute() > task.schedule.minute &&
        now.Second() > task.schedule.second
}

func (task *Task) FatherTasksAllDone() bool {
    fatherTasks := gjson.Parse(task.FatherTask).Array()
    for _, taskID := range fatherTasks {
        fatherTask := new(Task)
        fatherTask.LoadByKey(taskID.Uint())
        if !fatherTask.SuccessToday() {
            return false
        }
    }
    return true
}

func (task *Task) ShouldScheduleNow() bool {
    return task.ShouldScheduleToday() && task.ReachScheduleClock() && task.FatherTasksAllDone()
}

func (task *Task) NoJobToday() (bool) {
    jobs := []Job{}
    Fill(&jobs).Where("task_id = ? and date(created_at) = ? ", task.ID, datetime.Today())
    return len(jobs) == 0
}

func (task *Task) CreateJob() {
    job := new(Job)
    job.TaskID = task.ID
    job.Status = "pooling"
    job.DumpToMySQL()
}

func (task *Task) SuccessToday() bool {
    if !task.ShouldScheduleNow() || task.NoJobToday() {
        return false
    }
    job := new(Job)
    job.LoadByWhere("task_id = ? and date(created_at) = ? ", task.ID, datetime.Today())
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

func (task *Task) BeforeSave(scope *gorm.Scope) error {
    task.Schedule = fmt.Sprintf("%s %d %d-%d-%d %d:%d:%d", task.schedule.period, task.schedule.weekday,
        task.schedule.year, task.schedule.month, task.schedule.day,
        task.schedule.hour, task.schedule.minute, task.schedule.second,
    )
    return nil
}

func (task *Task) AfterSave(scope *gorm.Scope) error {
    values := strings.Fields(task.Schedule)
    period, weekday, date, clock := values[0], number.Int(values[1]), strings.Split(values[2], '-'), strings.Split(values[3], ':')
    year, month, day := number.Int(date[0]), number.Int(date[1]), number.Int(date[2])
    hour, minute, second := number.Int(clock[0]), number.Int(clock[1]), number.Int(clock[2])
    task.schedule = scheduleFormat{
        period: period,
        weekday:weekday,
        year: year,
        month: month,
        day:day,
        hour: hour,
        minute: minute,
        second:second,
    }
    return nil
}

