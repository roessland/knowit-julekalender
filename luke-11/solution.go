package main

import (
	"fmt"
	"time"
)

func main() {
	utc := time.FixedZone("UTC", 0)
	t := time.Date(1970, 1, 1, 0, 0, 0, 0, utc)
	unix64 := int64(0)
	unix32 := int32(0)
	extraHours := int64(0)
	for {
		t = t.Add(time.Duration(1) * time.Second)
		if t.Hour() == 23 && t.Minute() == 59 && t.Second() == 59 {
			extraHours += 1
		}
		unix64 = t.Unix() + extraHours*3600
		unix32 = int32(unix64)
		if unix64 != int64(unix32) {
			fmt.Println("something happened")
			fmt.Println(t.Add(time.Duration(-1) * time.Second))
			break
		}
	}
}
