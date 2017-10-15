package main

import (
    "fmt"
    "os"
    "./core/scheduler"
    "./core/worker"
    "./server"
    "./database"
)



func main() {
    if len(os.Args) != 2 {
        fmt.Println("必须附带启动参数：scheduler / worker / webserver / init")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "scheduler":
        fmt.Println("启动 scheduler")
        scheduler.Serve()
    case "worker":
        fmt.Println("启动 worker")
        worker.Serve()
    case "server":
        fmt.Println("启动 api 服务")
        server.Serve()
    case "init_db":
        database.Init()
    default:
        fmt.Println("不支持启动命令: " + os.Args[1])
    }
}
