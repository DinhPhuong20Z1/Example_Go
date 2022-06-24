package main

import (
	"fmt"
	"os"
)

func main() {

	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// Trì hoãn(defer) là một khái niệm khá mới trong điều khiển luồng. Nó cho phép một câu lệnh được gọi ra nhưng không được thực thi ngay mà hoãn lại đến khi các lệnh xung quanh trả về kết quả
// Câu lệnh được gọi qua từ khóa defer sẽ được đưa vào một stack, tức hoạt động theo cơ chế vào sau ra trước(last-in-first-out)
// lệnh nào defer sau sẽ được thực thi trước giốn như xếp 1 chồng đĩa thì chiếc đĩa sau cùng (ở trên cùng) sẽ được lấy ra trước.
