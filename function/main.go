package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	// giá trị res được gán bằn fuc plus sử dụng :=
	fmt.Println("res = ", res)

	res = plusPlus(1, 2, 3)
	// giá trị res được gán lại thì ko dùng := mà dùng =
	fmt.Println("res = ", res)
}
