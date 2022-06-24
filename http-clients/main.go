package main

import (
	"bufio"
	"fmt"
	"net/http"
)

// bắn data về server

func main() {
	resp, err := http.Get("https://gobyexample.com")

	// gọi api
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	// in ra dữ liệu 5 dòng đầu

	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
