package main

import "fmt"

func SlideLengthAndCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlide(s)

	s = s[:0]
	printSlide(s)

	s = s[:4]
	printSlide(s)

	s = s[2:]
	printSlide(s)
}

func printSlide(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}