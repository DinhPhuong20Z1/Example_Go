package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"b", "a", "c"}

	sort.Strings(strs)
	// sắp xếp theo string sort.Strings()

	fmt.Println("Strings", strs)

	ints := []int{10, 5, 50}

	sort.Ints(ints)
	// sắp xếp theo số nummber sort.Ints()
	fmt.Println("Ints:    ", ints)

	s := sort.IntsAreSorted(ints)
	// nếu trong mảng toàn là số là true
	fmt.Println("Sorted: ", s)
}
