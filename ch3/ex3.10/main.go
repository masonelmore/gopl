package main

import (
	"bytes"
	"fmt"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	// Figure out where to put the first comma.
	i := len(s) % 3
	if i == 0 {
		i = 3
	}

	var buf bytes.Buffer
	for len(s) > 3 {
		// Insert the digits before the comma.
		buf.WriteString(s[:i])
		buf.WriteString(",")
		// Chop off the digits we inserted.
		s = s[i:]
		// Reset the distance to the next comma.  Useful when we start with 1 or
		// 2 digits before the first comma.
		i = 3
	}

	// Insert the rest of the digits.
	buf.WriteString(s)

	return buf.String()
}

func main() {
	fmt.Println(comma("1"))
	fmt.Println(comma("12"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("12345678"))
	fmt.Println(comma("123456789"))
	fmt.Println(comma("1234567890"))
}
