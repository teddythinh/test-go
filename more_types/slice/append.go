package main

import "fmt"

func Append() {
	var s[]int
	PrintSlice2(s)

	// append works on nil slices
	s = append(s, 0)
	PrintSlice2(s)

	// the slice grows as needed
	s = append(s, 1)
	PrintSlice2(s)

	// we can add more than one element at a time
	s = append(s, 2, 3, 4)
	PrintSlice2(s)
}

func PrintSlice2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
