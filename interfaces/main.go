package main

import (
	"fmt"
	"math"
)

// định nghĩa chung về interface trong thế giới hướng đối tượng là "interface xác định hành vi của đối tượng".
// Nó chỉ xác định những gì đối tượng làm. Cách để thực hiện các hành vi tùy vào nó

type geometry interface {
	area() float64
	perim() float64
}

type total interface {
	totalTrue() float64
}

type rect struct {
	width, height float64
}

type cricle struct {
	radius float64
}

type string1 struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c cricle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c cricle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (w string1) totalTrue() float64 {
	return w.height + w.width
}

// func medusa(w total) {
// 	fmt.Println(w)
// 	println("w", w.totalTrue())
// }

func medusa(w total) {
	fmt.Println("w", w)
	fmt.Println("q", w.totalTrue())
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func PrintAll(vals []interface{}) {

	fmt.Println("vals123", vals)
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	// check nếu là rect thì mới nhảy vào hàm rect
	c := cricle{radius: 5}
	// check nếu là cricle thì mới nhảy vào cricle

	w := string1{width: 3, height: 4}

	medusa(w)

	measure(r)
	measure(c)

	names := []string{"Xevic", "Diana", "Join"}
	vals := make([]interface{}, len(names))
	for i, v := range names {

		vals[i] = v
	}

	PrintAll(vals)
}
