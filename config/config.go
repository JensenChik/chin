package config

import (
    "io/ioutil"
    "encoding/json"
    "github.com/sdbaiguanghe/glog"
    "path/filepath"
    "runtime"
)

type config struct {
    Core map[string]string
}

func loadConfig() (config) {
    var conf config
    _, thisFileAbsName, _, _ := runtime.Caller(0)
    config_path, _ := filepath.Abs(filepath.Dir(thisFileAbsName))
    filename := filepath.Join(config_path, "..", "chin.json")

    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        glog.Fatal("读取配置文件失败: ", err.Error())
    }
    if err := json.Unmarshal(bytes, &conf); err != nil {
        glog.Fatal("解析配置文件失败: ", err.Error())
    }
    return conf
}

var conf = loadConfig()
var SQL_CONN = conf.Core["sql_conn"]
var ROOT_NAME = conf.Core["root_name"]
var ROOT_PASSWD = conf.Core["root_passwd"]
var ROOT_MAIL = conf.Core["root_mail"]
var SECRET_KEY = []byte(conf.Core["secret_key"])
