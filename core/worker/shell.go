package worker

import (
    "os/exec"
    "github.com/sdbaiguanghe/glog"
    "io/ioutil"
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
        stdout, _ := ioutil.ReadAll(stdoutPipe)
        stderr, _ := ioutil.ReadAll(stderrPipe)
        sh.terminal.Wait()
        sh.Finish = true
        sh.Success = sh.terminal.ProcessState.Success()
        sh.Output = string(stdout) + string(stderr)
    }
    go fork()
}

func (sh *Shell) Kill() error {
    err := sh.terminal.Process.Kill()
    if err != nil {
        glog.Error("kill任务失败", err.Error())
    }
    sh.Output = "任务被手动杀死，当前系统无法保留过程输出"
    sh.Finish = true
    sh.Success = false
    return err
}


