package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())

	// thời gian thức tính từ 1970
	fmt.Println(now.UnixMilli())
	// tính theo giây

	fmt.Println(now.UnixNano())

	fmt.Println(time.Unix(now.Unix(), 0))

	fmt.Println(time.Unix(0, now.UnixNano()))

	// kỈ nguyên unix  Kỷ nguyên Unix là 00:00:00 UTC vào ngày 1 tháng 1 năm 1970 (một ngày tùy ý)
}
