package server

import (
    "net/http"
    "fmt"
)

func new_task(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "新建任务")
}
