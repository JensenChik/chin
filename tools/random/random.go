package random

import (
    "time"
    "math/rand"
)

func String(size int) string {
    str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    bytes := []byte(str)
    result := []byte{}
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i := 0; i < size; i++ {
        result = append(result, bytes[r.Intn(len(bytes))])
    }
    return string(result)
}

func Int(upbound int) uint {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    return uint(r.Intn(upbound))
}

func Float() float64 {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    return r.Float64()
}
