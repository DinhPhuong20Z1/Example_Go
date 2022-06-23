// Goroutine là một hàm có thể chạy đồng thời các hàm khác

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// khi một goroutine mới khởi chạy, tính năng goroutibe sẽ gọi trở lại ngay lập tức,
// không giống hàm, hệ thống không chờ Goroutine chạy xong mà trả về lập tức dòng code tiếp theo

// Goroutine chính cần chạy được để các goroutine khác chạy được,
//
func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	// cách fix deadline 1s
	fmt.Println("done")
}
