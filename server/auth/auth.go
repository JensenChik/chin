package auth

import (
    "fmt"
    "net/http"
    "github.com/gorilla/sessions"
    "../../database"
    "../../config"
    "github.com/sdbaiguanghe/glog"
)

var store = sessions.NewCookieStore(config.SECRET_KEY)
var sessionName = "sess"

func Login(w http.ResponseWriter, r *http.Request) {
    var userName = r.PostFormValue("user_name")
    var password = r.PostFormValue("password")
    if database.ExistsUser(userName, password) {
        glog.Debug("用户存在")
        session, err := store.Get(r, sessionName)
        if err != nil {
            glog.Error("获取session失败: ", err.Error())
        }
        session.Values["user_name"] = userName
        session.Save(r, w)
        fmt.Fprint(w, "用户存在收到登陆请求:" + userName + "@" + password)
    } else {
        //报异常
        fmt.Fprint(w, "用户不存在收到登陆请求:" + userName + "@" + password)
    }

}

func CurrentUser(r *http.Request) string {
    session, err := store.Get(r, sessionName)
    userName := session.Values["user_name"]
    if err != nil {
        glog.Error("获取session失败: ", err.Error())
        return ""
    } else if userName == nil {
        return ""
    } else {
        return userName.(string)
    }
}

func Logout(w http.ResponseWriter, r *http.Request) {
    glog.Debug("请求登出")
    fmt.Fprint(w, "请求登出")
}



