package main

import "fmt"

func fact(n int) int {

	// khai báo nhận tham số thì nó sẽ nhận từ 0 đến số đó
	if n == 0 {
		return 1
	}
	fmt.Println("n", n)
	return n * fact(n-1)

	// đệ quy nhân với nhau với các phần tử  đến số cuối cùng được gọi vào
}

func main() {
	fmt.Println(fact(3))
}
