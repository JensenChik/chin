package main

import (
    "os"
    "./core/scheduler"
    "./core/worker"
    "./server"
    "./database"
    "flag"
    "github.com/sdbaiguanghe/glog"
    "log"
)

func setDebug(ok bool) {
    flag.Lookup("logtostderr").Value.Set("true")
}

func main() {
    flag.String("as", "", "执行命令: scheduler / worker / webserver / init_db / mock_db")
    flag.Parse()
    setDebug(true)
    glog.SetLevelString("DEBUG")

    if log_dir := flag.Lookup("log_dir").Value.String(); log_dir != "" {
        log.Print("日志将保存到路径: ", log_dir)
    } else {
        log.Print("日志将保存到路径: ", os.TempDir())
    }
    defer glog.Flush()

    as := flag.Lookup("as").Value.String()
    if as == "" {
        glog.Fatal("必须指定as参数")
    }

    switch as {
    case "scheduler":
        glog.Info("启动 scheduler")
        scheduler.Serve()
    case "worker":
        glog.Info("启动 worker")
        worker.Serve()
    case "server":
        glog.Info("启动 api 服务")
        server.Serve()
    case "init_db":
        database.Init()
    case "mock_db" :
        database.Mock()
    default:
        glog.Fatal("不支持启动命令: " + as)
    }
}
