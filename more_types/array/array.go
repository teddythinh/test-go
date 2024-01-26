package main

import "fmt"

func MyArray() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1]) // Hello World
	fmt.Println(a) // [Hello World]

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4] // Slice
	fmt.Println(primes)
	fmt.Println(s)
}