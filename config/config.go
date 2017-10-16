package config

import (
    "log"
    "io/ioutil"
    "encoding/json"
    "os"
)

type config struct {
    Core map[string]string
}

func loadConfig() (config) {
    var conf config
    filename := "chin.json"
    if _, err:= os.Stat(filename); os.IsNotExist(err){
        os.Chdir("..")
    }

    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal("读取配置文件失败: ", err.Error())
    }
    if err := json.Unmarshal(bytes, &conf); err != nil {
        log.Fatal("解析配置文件失败: ", err.Error())
    }
    return conf
}

var conf = loadConfig()
var SQL_CONN = conf.Core["sql_conn"]
var ROOT_NAME = conf.Core["root_name"]
var ROOT_PASSWD = conf.Core["root_passwd"]
var ROOT_MAIL = conf.Core["root_mail"]
