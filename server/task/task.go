package task

import (
    "net/http"
    "fmt"
    "../auth"
)

func New(w http.ResponseWriter, r *http.Request) {
    currentUser := auth.CurrentUser(r)
    fmt.Fprint(w, "新建任务" + currentUser)
}
