package strftime

import (
	"time"
	"fmt"
)

func ExampleFormat() {
	t := time.Unix(1340244776, 0)
	utc, _ := time.LoadLocation("UTC")
	t = t.In(utc)
	fmt.Println(Format("%Y-%m-%d %H:%M:%S", t))
	// Output:
	// 2012-06-21 02:12:56
}