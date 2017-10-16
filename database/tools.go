package database

import (
    "time"
    "crypto/md5"
    "encoding/hex"
    "math/rand"
)

func toMD5(raw string) string {
    ctx := md5.New()
    ctx.Write([]byte(raw))
    return hex.EncodeToString(ctx.Sum(nil))
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

