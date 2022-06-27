package main

import "fmt"

// func hello(name string) string {}
// hello -> function name
// string in "name string" -> parameter type (loại tham số)
// string last -> return value (trả lại kết quả )

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func demoProps(length, width float64) (float64, float64) {
	var sup = length + width
	var total = length - width
	return sup, total
}

func rectPerimeter(a int, b int) int {
	return (a + b) * 2
}

func rectArea(a, b int) int {
	return a * b
}

func rectPeriAndArea2(a, b int) (rectPeria int, rectAr int) {
	rectPeria = rectPerimeter(a, b)
	rectAr = rectArea(a, b)
	return
}

// trong func () đầu tiên là giá trị mà func nhận vào
// () thứ 2 là giá trị trả ra khi thực thi gì đấy, có thể là một phép tính

func main() {
	res := plus(1, 2)
	// giá trị res được gán bằn fuc plus sử dụng :=
	fmt.Println("res = ", res)

	res = plusPlus(1, 2, 3)
	// giá trị res được gán lại thì ko dùng := mà dùng =
	fmt.Println("res = ", res)

	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %d Perimeter %d", area, perimeter)

	sup, _ := demoProps(90, 6)
	fmt.Println("total", sup)

	var num1 int = 2
	var num2 int = 4

	fmt.Println(num1, num2)

	a := 2
	b := 3

	rectPeri2, rectAr2 := rectPeriAndArea2(a, b)
	fmt.Println("rectPeri2", rectPeri2)
	fmt.Println("rectAr2", rectAr2)
	fmt.Printf("rectPeri2=%d, rectAr2=%d\n", rectPeri2, rectAr2)
	name := "Phuong"

	fmt.Printf("Xin chào,% v. Chào mừng!", name)

}
