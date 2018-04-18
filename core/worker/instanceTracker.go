package worker

import (
    "time"

    "../../model"
    "github.com/sdbaiguanghe/glog"
)

func registerMachineIfNotExists(stat *sysStat) uint {
    machine, _ := new(model.Machine).LoadByWhere("mac = ?", stat.MACAddress)
    if machine == nil {
        newMachine := model.Machine{
            MachineName: stat.HostName,
            IP:          stat.IP,
            MAC:         stat.MACAddress,
            CPULoad:     stat.Load5,
            MemoryLoad:  stat.MemUsedPercent,
            Alive:       true,
        }
        newMachine.DumpToMySQL()
    }
    return machine.ID
}

func instanceTracker() {
    glog.Info("instance tracker 开始启动")
    for {
        stat := getSysStat()
        machineID := registerMachineIfNotExists(stat)

        instances := []model.Instance{}
        model.Fill(&instances).Where("machine_id = ?", machineID)
        for _, instance := range instances {
            if instance.GetReady() {
                instance.MachineID = machineID
                instance.CreateAndRunShell()
            }
            glog.Error(instance.ID)
        }

        time.Sleep(time.Second)
    }
}
