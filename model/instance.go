package model

import (
    "github.com/jinzhu/gorm"
    "github.com/sdbaiguanghe/glog"
    "os/exec"
    "io"
    "bufio"
    "../tools/compress"
)

type Shell struct {
    Command  string
    Finish   bool
    Success  bool
    Output   string
    terminal *exec.Cmd
}

func (sh *Shell) Run() {
    sh.terminal = exec.Command("sh", "-c", sh.Command)

    fork := func() {
        stdoutPipe, _ := sh.terminal.StdoutPipe()
        stderrPipe, _ := sh.terminal.StderrPipe()
        sh.terminal.Start()
        stdoutReader := bufio.NewReader(stdoutPipe)
        stderrReader := bufio.NewReader(stderrPipe)
        for {
            line, err := stdoutReader.ReadString('\n')
            if err != nil || err == io.EOF {
                break
            }
            sh.Output += line
        }
        for {
            line, err := stderrReader.ReadString('\n')
            if err != nil || err == io.EOF {
                break
            }
            sh.Output += line
        }
        sh.terminal.Wait()
        sh.Finish = true
        sh.Success = sh.terminal.ProcessState.Success()
    }
    go fork()
}

func (sh *Shell) Kill() error {
    err := sh.terminal.Process.Kill()
    if err != nil {
        glog.Error("kill任务失败", err.Error())
    }
    sh.Finish = true
    sh.Success = false
    return err
}

type Instance struct {
    gorm.Model
    JobID     uint `gorm:"index"`
    MachineID uint
    StdOut    string `gorm:"type:longblob"`
    shell     Shell
}

func (Instance *Instance) GetReady() bool {
    return false
}

func (Instance *Instance) CreateAndRunShell() {

}


func (instance *Instance) BeforeSave(scope *gorm.Scope) error {
    instance.StdOut = compress.Zip(instance.StdOut)
    return nil
}

func (instance *Instance) AfterSave(scope *gorm.Scope) error {
    instance.StdOut = compress.Unzip(instance.StdOut)
    return nil
}

func (instance *Instance) AfterFind(scope *gorm.Scope) error {
    instance.StdOut = compress.Unzip(instance.StdOut)
    return nil
}

func (instance *Instance) DumpToMySQL() (bool, error) {
    ok, err := dumpToMysql(instance)
    return ok, err
}

func (instance *Instance) LoadByWhere(filters ...interface{}) (*Instance, error) {
    initInstance, err := loadByWhere(instance, filters...)
    if err != nil {
        return nil, err
    } else {
        return initInstance.(*Instance), nil
    }
}

func (instance *Instance) LoadByKey(key interface{}) (*Instance, error) {
    initInstance, err := loadByKey(instance, key)
    if err != nil {
        return nil, err
    } else {
        return initInstance.(*Instance), nil
    }
}

