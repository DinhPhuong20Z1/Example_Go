package main

import "fmt"

func main() {
	//ch := make(chan type, capacity)
	messages := make(chan string, 2)
	// Channels không có buffer sẽ làm cho goroutine tương ứng bị block
	// capacity cho biết số lương buffer  của channel được tạo ra. capacity của channel ko có buffer là bằng 0 thường được bỏ qua khi tạo channel
	messages <- "bufferred"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
