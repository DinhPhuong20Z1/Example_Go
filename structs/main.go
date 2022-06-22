// struct là một kiểu dữ liệu, nhưng ko giống các kiểu dữ liệu nguyên thủy thông thường như int, string...
// mà struct là kiểu dữ liệu người dùng tự định nghĩa
// một struct có thể bao gồm nhiều trường có các kiểu dữ liệu khác nhau

package main

import "fmt"

// Để định nghĩa một struct thì chúng ta dùng từ khóa type, từ khóa này để báo cho Go biết là chúng ta định nghĩa một struct
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	// truyền name vào
	p := person{name: name}
	//  nhận name  bắn vào  person
	p.age = 42
	// nhận thêm age bắn vào person
	return &p
	// trả kết quả bắn vào gán p
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Anlice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}

	fmt.Println(s.name)
	// gán rồi nhưng chỉ lấy name

	sp := &s
	// gán bằng ip s
	fmt.Println(sp.age)
	// có data  giống s nhưng chỉ lấy age

}
