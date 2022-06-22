package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()

	fmt.Println(a)
	fmt.Println(b)

	c, _ := vals()
	// Kí hiệu: _ (blank identifier) : bỏ qua giá trị mà ko có nhu cầu sử dụng
	fmt.Println(c)
}
