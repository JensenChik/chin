package number

import (
    "strconv"
)

func Uint(v interface{}) uint64 {
    switch v.(type) {
    case string:
        return str2uint(v.(string))
    }
    return uint64(0)
}

func str2uint(str string) uint64 {
    value, err := strconv.ParseUint(str, 10, 64)
    if err != nil {
        value = uint64(0)
    }
    return value
}

