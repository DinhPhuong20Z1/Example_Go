package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	// đợi 2s thì timer1 sẽ được kích hoạt

	<-timer1.C

	fmt.Println("Timer 1 fired")

	time2 := time.NewTimer(time.Second)

	go func() {
		<-time2.C
		fmt.Println("timer 2 fired")
	}()

	stop2 := time2.Stop()
	// Hủy lệnh đợi của sleep
	fmt.Println("stop2", stop2)
	if stop2 {
		fmt.Println("timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
	// ko cần đợi sau khi timer1 kích hoạt thì sẽ chạy luôn
}
