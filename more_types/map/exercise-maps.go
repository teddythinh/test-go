package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	// Create a map of string to int
	m := make(map[string]int)
	// Split the string into a slice of strings
	words := strings.Fields(s)
	// Iterate over the slice of strings
	for _, word := range words {
		// Increment the count of the word
		m[word]++
	}
	// Return the map
	return m
}

func ExerciseMaps() {
	wc.Test(WordCount)
}