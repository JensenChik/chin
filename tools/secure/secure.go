package secure

import (
    "crypto/md5"
    "encoding/hex"
)

func MD5(raw string) string {
    ctx := md5.New()
    ctx.Write([]byte(raw))
    return hex.EncodeToString(ctx.Sum(nil))
}



