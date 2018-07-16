package main

import (
	"bytes"
	"strings"
)

func main() {
	comma("1.123")
}

func comma(s string) string {
	var buf bytes.Buffer

	// Strip the sign from s and put in the buffer.
	if strings.HasPrefix(s, "-") {
		buf.WriteString("-")
		s = s[1:]
	}

	// Strip the decimal portion from s and save for the end.
	var dec string
	if dot := strings.Index(s, "."); dot != -1 {
		dec = s[dot:]
		s = s[:dot]
	}

	// Commafy the remaining integer part of s.
	i := len(s) % 3
	if i == 0 {
		i = 3
	}
	for len(s) > 3 {
		buf.WriteString(s[:i])
		buf.WriteString(",")
		s = s[i:]
		i = 3
	}
	buf.WriteString(s)

	buf.WriteString(dec)
	return buf.String()
}
