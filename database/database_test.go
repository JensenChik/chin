package database

import (
    "testing"
    "log"
    "os"
)

func TestMain(m *testing.M) {
    log.Print("setup: 初始化数据库"); Init()
    returnCode := m.Run()
    log.Print("teardown: 清理数据库"); Init()
    os.Exit(returnCode)
}
