package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// chạy đồng thời 2 func

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			// có kết quả 1 in ra
			fmt.Println("received", msg1)

		case msg2 := <-c2:
			// có kết quả 2 in ra
			fmt.Println("received", msg2)
		}
	}
}
