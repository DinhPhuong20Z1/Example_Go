package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	//nhận địa chỉ ip biến i
	*iptr = 0
	// gán giá trị vào ip i rồi trả ra
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)
	// chưa được gán lại giá trị vào i

	zeroptr(&i)
	// lấy địa chỉ ip của biến i
	// có địa chỉ ip của biến i truyền vào
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	//  địa chỉ ip của biến i

}
