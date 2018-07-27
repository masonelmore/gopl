package reverse

import (
	"unicode/utf8"
)

// reverse reverses the characters of a UTF8-encoded []byte slice in-place by
// reversing all multi-byte runes before reversing the entire []byte slice.
func reverse(bytes []byte) []byte {
	if len(bytes) < 2 {
		return bytes
	}

	for i := 0; i < len(bytes); {
		_, n := utf8.DecodeRune(bytes[i:])
		if n > 1 {
			rev(bytes[i : i+n])
		}
		i += n
	}

	rev(bytes)

	return bytes
}

// reverse2 does the same thing as reverse except it reverses the entire []byte
// slice and then fixes all multi-byte runes by reversing them back to normal.
func reverse2(bytes []byte) []byte {
	if len(bytes) < 2 {
		return bytes
	}

	rev(bytes)

	for i, runeEnd := 0, 0; i < len(bytes); i++ {
		if utf8.RuneStart(bytes[i]) {
			if i-runeEnd > 1 {
				rev(bytes[runeEnd : i+1])
			}
			runeEnd = i + 1
		}
	}

	return bytes
}

func rev(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
