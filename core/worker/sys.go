package worker

import (
    "strconv"
    "strings"
    "io/ioutil"
    "bufio"
    "os"
    "net"
    "github.com/sdbaiguanghe/glog"
    "reflect"
)

type sysStat struct {
    HostName    string  // 机器名
    IP          string  // IP地址
    MACAddress  string  // MAC地址

                        // /proc/loadavg
    Load1       float64 // 1分钟平均负载
    Load5       float64 // 5分钟平均负载
    Load15      float64 // 15分钟平均负载

                        // /proc/meminfo


    MemTotal       uint64  // 内存总量
    MemFree        uint64  // 未使用内存总量
    MemActive      uint64  // 最近经常被使用的内存大小总量
    MemInactive    uint64  // 最近不是经常使用的内存
    MemBuffer     uint64  // 临时存储原始磁盘块的总量
    MemCache      uint64  // 用作缓存内存的物理内存总量
    Available   uint64  // 可使用内存总量
    Used        uint64  // 已使用内存总量
    UsedPercent float64 // 已使用百分比
    Wired       uint64  // 联动内存总量(mac os || BSD)
    Shared      uint64  // 多个进程共享的内存总量

                        // /proc/net/dev
                        // sum of all interfaces
    Name        string  // 网卡名
    BytesSent   uint64  // 发送字节数
    BytesRecv   uint64  // 接受字节数
    PacketsSent uint64  // 发送包数
    PacketsRecv uint64  // 接收包数

                        // /proc/stat
    CPU         string  //CPU名称
    User        float64 // 用户
    System      float64 //
    Idle        float64 // 除硬盘IO等待时间以外其它等待时间
    Nice        float64 // nice值为负的进程所占用的CPU时间
    Iowait      float64 // 硬盘IO等待时间
    Irq         float64 // 硬中断时间
    Softirq     float64 // 软中断时间
    Steal       float64
    Guest       float64
    GuestNice   float64
    Stolen      float64
}

func getSysStat() *sysStat {
    stat := new(sysStat)
    stat.cpuTimes()
    stat.netIOCounters()
    stat.virtualMemory()
    stat.loadAvg()
    return stat
}

func readLinesBetween(filename string, offset uint, n int) []string {
    f, err := os.Open(filename)
    if err != nil {
        glog.Fatal(err, "读取文件失败")
    }
    defer f.Close()

    var lines []string
    r := bufio.NewReader(f)
    for i := 0; i < n + int(offset) || n < 0; i++ {
        line, err := r.ReadString('\n')
        if err != nil {
            break
        }
        if i < int(offset) {
            continue
        }
        lines = append(lines, strings.Trim(line, "\n"))
    }

    return lines
}

func readLines(filename string) []string {
    return readLinesBetween(filename, 0, -1)
}

func getLocalIP() string {
    addresses, err := net.InterfaceAddrs()
    if err != nil {
        glog.Fatal(err, "获取机器 IP 失败")
    }
    for _, address := range addresses {
        if IPNet, ok := address.(*net.IPNet); ok && !IPNet.IP.IsLoopback() {
            if IPNet.IP.To4() != nil {
                return IPNet.IP.String()
            }

        }
    }
    return "loaclhost"
}

func getMacAddress() string {
    eth0, err := net.InterfaceByName("eth0")
    if err != nil {
        glog.Fatal(err, "该机器没有 eth0 网卡，无法获取MAC地址")
    }
    return eth0.HardwareAddr.String()
}

func (stat *sysStat) cpuTimes() {
    filename := "/proc/stat"
    lines := readLinesBetween(filename, 0, 1)
    stat.parseStatLine(lines[0])
}

func (stat *sysStat) parseStatLine(line string) error {
    fields := strings.Fields(line)

    if strings.HasPrefix(fields[0], "cpu") == false {
        glog.Fatal("not contain cpu")
    }

    cpu := fields[0]
    if cpu == "cpu" {
        cpu = "cpu-total"
    }
    user, err := strconv.ParseFloat(fields[1], 64)
    if err != nil {
        return err
    }
    nice, err := strconv.ParseFloat(fields[2], 64)
    if err != nil {
        return err
    }
    system, err := strconv.ParseFloat(fields[3], 64)
    if err != nil {
        return err
    }
    idle, err := strconv.ParseFloat(fields[4], 64)
    if err != nil {
        return err
    }
    iowait, err := strconv.ParseFloat(fields[5], 64)
    if err != nil {
        return err
    }
    irq, err := strconv.ParseFloat(fields[6], 64)
    if err != nil {
        return err
    }
    softirq, err := strconv.ParseFloat(fields[7], 64)
    if err != nil {
        return err
    }
    stolen, err := strconv.ParseFloat(fields[8], 64)
    if err != nil {
        return err
    }

    cpu_tick := float64(100) // TODO: how to get _SC_CLK_TCK ?

    stat.CPU = cpu
    stat.User = float64(user) / cpu_tick
    stat.Nice = float64(nice) / cpu_tick
    stat.System = float64(system) / cpu_tick
    stat.Idle = float64(idle) / cpu_tick
    stat.Iowait = float64(iowait) / cpu_tick
    stat.Irq = float64(irq) / cpu_tick
    stat.Softirq = float64(softirq) / cpu_tick
    stat.Stolen = float64(stolen) / cpu_tick

    if len(fields) > 9 {
        // Linux >= 2.6.11
        steal, err := strconv.ParseFloat(fields[9], 64)
        if err != nil {
            return err
        }
        stat.Steal = float64(steal)
    }
    if len(fields) > 10 {
        // Linux >= 2.6.24
        guest, err := strconv.ParseFloat(fields[10], 64)
        if err != nil {
            return err
        }
        stat.Guest = float64(guest)
    }
    if len(fields) > 11 {
        // Linux >= 3.2.0
        guestNice, err := strconv.ParseFloat(fields[11], 64)
        if err != nil {
            return err
        }
        stat.GuestNice = float64(guestNice)
    }

    return nil
}

type netIOCountersStat struct {
    Name        string
    BytesSent   uint64
    BytesRecv   uint64
    PacketsSent uint64
    PacketsRecv uint64
}

func (stat *sysStat) netIOCounters() error {
    filename := "/proc/net/dev"
    lines := readLines(filename)
    statlen := len(lines) - 1

    all := make([]netIOCountersStat, 0, statlen)

    for _, line := range lines[2:] {
        parts := strings.SplitN(line, ":", 2)
        if len(parts) != 2 {
            continue
        }
        interfaceName := strings.TrimSpace(parts[0])
        if interfaceName == "" {
            continue
        }

        fields := strings.Fields(strings.TrimSpace(parts[1]))
        bytesRecv, err := strconv.ParseUint(fields[0], 10, 64)
        if err != nil {
            return err
        }
        packetsRecv, err := strconv.ParseUint(fields[1], 10, 64)
        if err != nil {
            return err
        }
        bytesSent, err := strconv.ParseUint(fields[8], 10, 64)
        if err != nil {
            return err
        }
        packetsSent, err := strconv.ParseUint(fields[9], 10, 64)
        if err != nil {
            return err
        }

        nic := netIOCountersStat{
            Name:        interfaceName,
            BytesRecv:   bytesRecv,
            PacketsRecv: packetsRecv,
            BytesSent:   bytesSent,
            PacketsSent: packetsSent}

        all = append(all, nic)
    }

    return stat.getNetIOCountersAll(all)
}

func (stat *sysStat) getNetIOCountersAll(n []netIOCountersStat) error {
    stat.Name = "all-interfaces"
    for _, nic := range n {
        stat.BytesRecv += nic.BytesRecv
        stat.PacketsRecv += nic.PacketsRecv
        stat.BytesSent += nic.BytesSent
        stat.PacketsSent += nic.PacketsSent
    }
    return nil
}

func (stat *sysStat) virtualMemory() {
    lines := readLines("/proc/meminfo")
    for _, line := range lines {
        kv := strings.Split(line, ":")
        fieldName, hit := map[string]string{
            "MemTotal": "MemTotal",
            "MemFree": "MemFree",
            "Buffers": "MemBuffer",
            "Cached": "MemCache",
            "Active": "MemActive",
            "Inactive": "MemInactive",
        }[strings.TrimSpace(kv[0])]
        if hit {
            value, _ := strconv.ParseUint(strings.Replace(strings.TrimSpace(kv[1]), " kB", "", -1), 10, 64)
            value /= 1000
            reflect.ValueOf(stat).Elem().FieldByName(fieldName).SetUint(value)
        }
    }
    stat.Available = stat.MemFree + stat.MemBuffer + stat.MemCache
    stat.Used = stat.MemTotal - stat.MemFree
    stat.UsedPercent = float64(stat.MemTotal - stat.Available) / float64(stat.MemTotal) * 100.0
}

func (stat *sysStat) loadAvg() {
    line, _ := ioutil.ReadFile("/proc/loadavg")
    values := strings.Fields(string(line))
    stat.Load1, _ = strconv.ParseFloat(values[0], 64)
    stat.Load5, _ = strconv.ParseFloat(values[1], 64)
    stat.Load15, _ = strconv.ParseFloat(values[2], 64)
}
