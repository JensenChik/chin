package database

import (
    "time"
    "crypto/md5"
    "encoding/hex"
    "math/rand"
    "bytes"
    "compress/zlib"
    "io"
)

func toMD5(raw string) string {
    ctx := md5.New()
    ctx.Write([]byte(raw))
    return hex.EncodeToString(ctx.Sum(nil))
}

func zip(raw string) string {
    var buffer bytes.Buffer
    writer := zlib.NewWriter(&buffer)
    writer.Write([]byte(raw))
    writer.Close()
    return buffer.String()
}

func unzip(raw string) string {
    reader := bytes.NewReader([]byte(raw))
    var buffer bytes.Buffer
    r, _ := zlib.NewReader(reader)
    io.Copy(&buffer, r)
    return buffer.String()
}

func randomString(size int) string {
    str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    bytes := []byte(str)
    result := []byte{}
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i := 0; i < size; i++ {
        result = append(result, bytes[r.Intn(len(bytes))])
    }
    return string(result)
}

