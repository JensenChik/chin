package datetime

import "time"

func Today() string {
    return time.Now().Format("2016-01-02")
}
