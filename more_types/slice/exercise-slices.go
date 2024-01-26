package main

func Pic(dx, dy int) [][]uint8 {
	// Create a slice of length dy
	s := make([][]uint8, dy)

	// For each row, create a slice of length dx
	for i := range s {
		s[i] = make([]uint8, dx)

		// For each column, set the value to the product of the row and column
		for j := range s[i] {
			s[i][j] = uint8(i * j)
		}
	}
	
	return s
}