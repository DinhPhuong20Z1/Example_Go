package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	// hiểu là gán hàm bằng 1 giá trị khác, mỗi khi gọi vào giá trị được gán thì nó sẽ gọi vòa hàm gán

	newInts := intSeq()

	fmt.Println(newInts())
	// khi thay đổi giá trị được gán thì nó sẽ làm mới
}
