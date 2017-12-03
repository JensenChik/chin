package database

import "github.com/jinzhu/gorm"

type Machine struct {
    gorm.Model
    MachineName string
    IP          string `gorm:"size:15"`
    MAC         string `gorm:"size:17"`
    CPULoad     int
    MemoryLoad  int
    Alive       bool
}

func (machine *Machine) DumpToMySQL() (bool, error) {
    ok, err := dumpToMysql(machine)
    return ok, err
}

func (machine *Machine) LoadByWhere(filters ...interface{}) (*Machine, error) {
    initMachine, err := loadByWhere(machine, filters...)
    if err != nil {
        return nil, err
    } else {
        return initMachine.(*Machine), nil
    }
}

func (machine *Machine) LoadByKey(key interface{}) (*Machine, error) {
    initMachine, err := loadByKey(machine, key)
    if err != nil {
        return nil, err
    } else {
        return initMachine.(*Machine), nil
    }
}
