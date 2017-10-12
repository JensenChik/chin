package server

import (
    "fmt"
    "net/http"
    "log"
)

func Serve() {
    fmt.Println("api服务器开始服务")
    http.HandleFunc("/login", login)
    http.HandleFunc("/logout", logout)
    http.HandleFunc("/new_task", new_task)
    log.Fatal(http.ListenAndServe("localhost:6421", nil))
}

