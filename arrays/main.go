package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println("emp", a)

	a[4] = 100

	fmt.Println("set", a)
	fmt.Println("get", a[4])

	fmt.Println("len", len(a))

	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println("dcl", b)

	var twoD [2][3]int

	fmt.Println("abc", twoD)
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d", twoD)

	var x [2]string

	x[0] = "Hello"
	x[1] = "World"
	fmt.Println("x", x)

	var z [3]complex128 // khai báo mảng kiểu complex128 3 phần tử
	fmt.Println(z)

	// Khởi tạo một mảng gồm 6 số kiểu int và gán luôn giá trị cho nó
	number := [6]int{2, 3, 5, 7, 11}
	fmt.Println(number)

	// khai báo mảng nhưng không ghi rõ kích thước,
	// trình biên dịch sẽ tự hiểu dựa vào số phần tử đã khai báo
	ar := [...]int{12, 78, 50}
	fmt.Println(ar)

	y := [4]float64{3.5, 7.2, 4.8, 9.5}
	sum := float64(0)

	for i := 0; i < len(y); i++ {
		sum = sum + y[i]
	}
	fmt.Printf("Sum of all the elements in array  %v = %f\n", y, sum)

	// mảng 2 chiều

	zlib := [2][2]int{
		{3, 5},
		{7, 9},
	}
	fmt.Println(zlib)

	w := [5][4]float64{
		{1, 3},
		{4.5, -3, 7.4, 2},
		{6, 2, 11},
	}
	fmt.Println(w)
}
