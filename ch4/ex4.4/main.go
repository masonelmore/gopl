package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	a = rotate(a, 3)
	// original: [0, 1, 2, 3, 4, 5]
	// rotated:  [3, 4, 5, 0, 1, 2]
	fmt.Println(a)
}

// rotate rotates a slice of ints.  Positive shift values rotate to the right,
// and negative shift values rotate to the left.
func rotate(ints []int, shift int) []int {
	r := make([]int, len(ints))

	if len(ints) <= 1 || shift == 0 {
		// We could just return ints here, but making a copy is more consistent
		// with the behavior of the rest of the function.
		copy(r, ints)
		return r
	}

	// TODO: Try to find a more clever way to do this without an if statement.
	if shift < 0 {
		shift *= -1
		for i := 0; i < len(ints); i++ {
			j := (i + shift) % len(ints)
			r[i] = ints[j]
		}
	} else {
		for i := 0; i < len(ints); i++ {
			j := (i + shift) % len(ints)
			r[j] = ints[i]
		}
	}

	return r
}
