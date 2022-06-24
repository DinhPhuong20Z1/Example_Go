package main

import "os"

func main() {
	panic("a problem")

	_, err := os.Create("./tmp/file")
	if err != nil {
		panic(err)
	}

}

// Pantic(hoảng loạn)
// Trong một vài trường hợp trương trình không thể tiếp tục chạy bình thường khi gặp tình huống bất thường.
// Trong trường hợp này chúng ta sử dụng Pantic để chấm dứt chương trình
// khi một hàm gặp Pantic , nó lập tức dừng xử lý, hàm bất kì được defer sẽ được chạy và sau đó kiểm soát trả về cho phía gọi

// https://vngeeks.com/go-panic-va-recover/
