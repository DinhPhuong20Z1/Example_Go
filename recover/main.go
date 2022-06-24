package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("Afer mayPanic()")

}

// Recover(khôi phục) là một hàm được tích hợp sẵn sàng trong Go, được sử dụng để lấy lại kiểm soát của goroutine đang panic
