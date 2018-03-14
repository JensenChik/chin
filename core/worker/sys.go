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
    OS             string
    HostName       string  // 机器名
    MACAddress     string  // MAC地址

    Load1          float64 // 1分钟平均负载
    Load5          float64 // 5分钟平均负载
    Load15         float64 // 15分钟平均负载

    MemTotal       uint64  // 内存总量
    MemFree        uint64  // 未使用内存总量
    MemActive      uint64  // 最近经常被使用的内存大小总量
    MemInactive    uint64  // 最近不是经常使用的内存
    MemBuffer      uint64  // 临时存储原始磁盘块的总量
    MemCache       uint64  // 用作缓存内存的物理内存总量
    MemAvailable   uint64  // 可使用内存总量
    MemUsed        uint64  // 已使用内存总量
    MemUsedPercent float64 // 已使用百分比

    IP             string  // IP地址
    NetSendByte    uint64  // 发送字节数
    NetSendPack    uint64  // 发送包数
    NetRecvByte    uint64  // 接受字节数
    NetRecvPack    uint64  // 接收包数
}

func getSysStat() *sysStat {
    stat := new(sysStat)
    stat.loadAvg()
    stat.virtualMemory()
    stat.network()
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

func firstLineOf(filename string) string {
    return strings.TrimSpace(readLinesBetween(filename, 0, 1)[0])
}

func str2uint(str string) uint64 {
    value, err := strconv.ParseUint(str, 10, 64)
    if err != nil {
        value = 0
    }
    return value
}

func (stat *sysStat)sysInfo() {
    stat.OS = strings.TrimSpace(strings.Split(firstLineOf("/etc/issue"), `\`)[0])
    stat.HostName = firstLineOf("/proc/sys/kernel/hostname")
    stat.MACAddress = firstLineOf("/sys/class/net/eth0/address")
}

func (stat *sysStat) loadAvg() {
    line, _ := ioutil.ReadFile("/proc/loadavg")
    values := strings.Fields(string(line))
    stat.Load1, _ = strconv.ParseFloat(values[0], 64)
    stat.Load5, _ = strconv.ParseFloat(values[1], 64)
    stat.Load15, _ = strconv.ParseFloat(values[2], 64)
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
            value := str2uint(strings.Replace(strings.TrimSpace(kv[1]), " kB", "", -1)) / 1000
            reflect.ValueOf(stat).Elem().FieldByName(fieldName).SetUint(value)
        }
    }
    stat.MemAvailable = stat.MemFree + stat.MemBuffer + stat.MemCache
    stat.MemUsed = stat.MemTotal - stat.MemFree
    stat.MemUsedPercent = float64(stat.MemTotal - stat.MemAvailable) / float64(stat.MemTotal) * 100.0
}

func (stat *sysStat) network() {
    addresses, _ := net.InterfaceAddrs()
    for _, address := range addresses {
        if IPNet, ok := address.(*net.IPNet); ok && !IPNet.IP.IsLoopback() && IPNet.IP.To4() != nil {
            stat.IP = IPNet.IP.String()
        }
    }
    for _, line := range readLines("/proc/net/dev")[2:] {
        values := strings.Fields(strings.TrimSpace(strings.SplitN(line, ":", 2)[1]))
        stat.NetRecvByte += str2uint(values[0]) / 1000000
        stat.NetRecvPack += str2uint(values[1])
        stat.NetSendByte += str2uint(values[8]) / 1000000
        stat.NetSendPack += str2uint(values[9])
    }
}


