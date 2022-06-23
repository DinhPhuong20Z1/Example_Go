// Channels có thể được coi là các đường ống sử dụng mà Goroutine  giao tiếp

package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	// dùng để giao tiếp
	msg := <-messages
	fmt.Println(msg)
}
