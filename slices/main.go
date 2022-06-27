package main

import "fmt"

// slice không có bất kì dữ liệu nào, chúng là  các tham chiếu đến mảng hiện có, nó mô tả một phần hoặc toàn bô ayrray
// nó mô tả một phần (hoặc toàn bộ)Array. Nó có kích thước rộng nên thường được sử dụng nhiều hơn array
// slices có thể tạo ra từ một Arrau bằng cách lấy  từ một vị trí phần tử bắt đầu và kết thúc trong array

func main() {
	s := make([]string, 3)
	fmt.Println("emp", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set", s)
	fmt.Println("get", s[2])

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[2:]
	fmt.Println("sl2:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	number := [6]int{2, 4, 5, 7, 8, 10}

	// Khởi tạo Slice s bằng cách cắt từ phần tử ở vị trí 2 đến phần tử ở vị trí 5 của Array number
	var w []int = number[2:5]

	// In ra giá trị của Slice s
	fmt.Println(w) // Giá trị của s là [5, 7, 8]

	data1 := make([]int, 5)

	printSlice("data1", data1)

}

func printSlice(s string, x []int) {
	fmt.Print("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
