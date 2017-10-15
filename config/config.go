package config

import (
    "log"
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type config struct {
    Core map[string]string
}

func load_config(filename string) (config) {
    var conf config
    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("Read file:", err.Error())
        log.Fatal("读取配置文件失败")
    }
    if err := json.Unmarshal(bytes, &conf); err != nil {
        fmt.Println("unmarshal:", err.Error())
        log.Fatal("解析配置文件失败")
    }
    return conf
}

var conf = load_config("config/chin.json")
var SQL_CONN = conf.Core["sql_conn"]
