package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	f := fibonacci()

	// Test the first 10 Fibonacci numbers
	expected := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	for i, want := range expected {
		got := f()
		if got != want {
			t.Errorf("fibonacci() returned %d, want %d at index %d", got, want, i)
		}
	}
}
