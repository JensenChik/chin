package compress

import (
    "bytes"
    "compress/zlib"
    "io"
)

func Zip(raw string) string {
    var buffer bytes.Buffer
    writer := zlib.NewWriter(&buffer)
    writer.Write([]byte(raw))
    writer.Close()
    return buffer.String()
}

func Unzip(raw string) string {
    reader := bytes.NewReader([]byte(raw))
    var buffer bytes.Buffer
    r, _ := zlib.NewReader(reader)
    io.Copy(&buffer, r)
    return buffer.String()
}
