// generics là một kĩ thuật lập trình thường xuất hiện trong OOP( lập trình hướng đối tương )
// kĩ thuật này giúp định nghĩa các hàm (phương thức method/function) chấp nhận các tham số chung

// Nhược điểm: khi trương trình đang thực thi(runtime) thì mới phát hiện lỗi, giảm tính an toàn của Go. Ảnh hưởng đến tốc độ thực thi(performance)
// ƯU điểm: tải sử dụng code,
package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	// tên biến dùng làm phần "giữ chỗ". Thật ra trong Go gọi là type (kiểu biến): VD K,V,...
	// ràng buộc (constraint) cho biển rõ ràng hơn, cụ thể hơn một chút. Nếu không có gì ràng buộc thì sẽ là any

	// K và V là 2 kiểu dữ liệu generics cho hàm MapKeys
	// cuối hàm MapKeys phải trả về một mảng mới []K
	r := make([]K, 0, len(m))

	for k := range m {
		r = append(r, k)
	}

	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next

	}
}

func (lst *List[T]) GetAll() []T {
	fmt.Println("abvc", lst)
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
		fmt.Println("lst.head", elems)
	}

	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	fmt.Println("key map:", MapKeys(m))

	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	fmt.Println("list:", lst.GetAll())

}
