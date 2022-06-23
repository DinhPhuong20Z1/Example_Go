package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(2000 * time.Millisecond)

	// đợi 2000ms thì sẽ bắn tần cái 1

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())

		//	sau khi bắn hết 5 lần thì sẽ bắn 3 lần tiếp theo
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(4000 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())

		// 2 lần cuối bắn cách nhau 4s
	}
}
