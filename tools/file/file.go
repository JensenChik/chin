package file

import (
    "bufio"
    "os"
    "strings"

    "github.com/sdbaiguanghe/glog"
)

func ReadLinesBetween(filename string, offset uint, n int) []string {
    f, err := os.Open(filename)
    if err != nil {
        glog.Fatal(err, "读取文件失败")
    }
    defer f.Close()

    var lines []string
    r := bufio.NewReader(f)
    for i := 0; i < n+int(offset) || n < 0; i++ {
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

func ReadLines(filename string) []string {
    return ReadLinesBetween(filename, 0, -1)
}

func FirstLineOf(filename string) string {
    return strings.TrimSpace(ReadLinesBetween(filename, 0, 1)[0])
}
