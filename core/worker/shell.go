package worker

import (
    "os/exec"
    "github.com/sdbaiguanghe/glog"
    "bufio"
    "io"
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


