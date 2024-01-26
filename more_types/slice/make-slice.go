package main

import "fmt"

func MakeSlices() {
	a := make([]int, 5)
	PrintSlice("a", a)

	b := make([]int, 0, 5)
	PrintSlice("b", b)

	c := b[:2]
	PrintSlice("c", c)

	d := c[2:5]
	PrintSlice("d", d)
}

func PrintSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}