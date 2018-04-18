package backend

import (
    "fmt"
    "log"
    "net/http"

    "./auth"
    "./task"
)

func Serve() {
    fmt.Println("api服务器开始服务")
    http.HandleFunc("/login", auth.Login)
    http.HandleFunc("/logout", auth.Logout)
    http.HandleFunc("/new_task", task.New)
    log.Fatal(http.ListenAndServe("localhost:6421", nil))
}
