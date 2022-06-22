package main

import "fmt"

func main() {
	var a = "initial"

	l := true  //a được gán bằng true
	g := false // b được gán bằng false

	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var e int
	fmt.Println(e)

	f := g && l

	fmt.Println(f)
}
