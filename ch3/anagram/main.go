// Anagram reports whether two strings contain the same letters in a different
// order.
package main

import (
	"strings"
)

func main() {}

func anagram(a, b string) bool {
	// Must not be identical
	if a == b {
		return false
	}

	// Must be the same length
	if len(a) != len(b) {
		return false
	}

	// Must have 2 or more characters
	if len(a) < 2 {
		return false
	}

	// Must have the same number of each character
	for _, c := range a {
		aCount := strings.Count(a, string(c))
		bCount := strings.Count(b, string(c))
		if aCount != bCount {
			return false
		}
	}

	// All tests passed.  They must be anagrams of each other.
	return true
}
