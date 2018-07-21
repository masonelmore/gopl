package reverse

import (
	"fmt"
	"testing"
)

var funcs = []struct {
	name string
	f    func([]byte) []byte
}{
	{"reverse", reverse},
	{"reverse2", reverse2},
}

func TestReverse(t *testing.T) {
	var tests = []struct {
		input, want string
	}{
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"abcd", "dcba"},
		{"abcde", "edcba"},
		{"😨", "😨"},
		{"a😨", "😨a"},
		{"a😨c", "c😨a"},
		{"𠜎", "𠜎"},
		{"𠜎😨", "😨𠜎"},
		{"𠜎😨a", "a😨𠜎"},
		{"𠜎😨aa😨𠜎", "𠜎😨aa😨𠜎"},
	}

	for _, f := range funcs {
		t.Run(f.name, func(t *testing.T) {
			for _, test := range tests {
				input := []byte(test.input)
				orig := make([]byte, len(input))
				copy(orig, input)
				want := []byte(test.want)
				if f.f(input); !eq(input, want) {
					t.Errorf("reverse(%q), got %x (), want %x", orig, input, want)
				}
			}
		})
	}
}

func eq(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkReverse(b *testing.B) {
	var benchmarks = []struct {
		name, input string
	}{
		{"ASCII short", "abcdef"},
		{"ASCII medium", "abcdefghijklmnopqrstuvwxyz"},
		{"ASCII long", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"},
		{"ASCII super-long", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"},
		{"UTF8 short", "𠜎𠜎😨😨"},
		{"UTF8 medium", "𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨"},
		{"UTF8 long", "𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨"},
		{"UTF8 super-long", "𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨𠜎𠜎😨😨"},
	}

	for _, f := range funcs {
		for _, benchmark := range benchmarks {
			name := fmt.Sprintf("%s %s", f.name, benchmark.name)
			input := []byte(benchmark.input)
			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f.f(input)
				}
			})
		}
	}
}
