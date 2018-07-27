package squash

import (
	"unicode"
	"unicode/utf8"
)

// squash squashes adjacent Unicode spaces into a single ASCII space. Any
// standalone rune that matches unicode.IsSpace will NOT be replaced with an
// ASCII space, per the exercise description.
func squash(bytes []byte) []byte {
	if len(bytes) < 2 {
		return bytes
	}

	out := bytes[:0]
	run := 0
	for i := 0; i < len(bytes); {
		thisRune, size := utf8.DecodeRune(bytes[i:])

		if i+size >= len(bytes) {
			if run > 0 {
				out = append(out, ' ')
			} else {
				out = append(out, bytes[i:]...)
			}
			return out
		}

		nextRune, _ := utf8.DecodeRune(bytes[i+size:])
		if unicode.IsSpace(thisRune) && unicode.IsSpace(nextRune) {
			run += size
		} else if run > 0 {
			out = append(out, ' ')
			run = 0
		} else {
			out = append(out, bytes[i:i+size]...)
		}

		i += size
	}

	return out
}
