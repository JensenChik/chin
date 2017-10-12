package server

import (
    "fmt"
    "net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "请求登陆")
}

func logout(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "请求登出")
}
