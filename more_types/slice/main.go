package main

import "golang.org/x/tour/pic"

func main() {

	// Slice
	SlicePointers()

	// Slice Literals
	SliceLiterals()

	// Slice defaults
	SlideDefaults()

	// Slice length and capacity
	SlideLengthAndCapacity()

	// Nil slices
	NilSlices()

	// Make slices
	MakeSlices()

	// Slices of slice
	SlicesOfSlice()

	// Append
	Append()

	// Exercise: Slices
	pic.Show(Pic)
}