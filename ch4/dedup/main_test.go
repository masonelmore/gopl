package main

import (
	"fmt"
	"testing"
)

var funcs = []struct {
	name string
	f    func([]string) []string
}{
	{"dupAppend", dupAppend},
	{"dupCopy", dupCopy},
	{"dupRunAppend", dupRunAppend},
	{"dupRunCopy", dupRunCopy},
}

func TestDup(t *testing.T) {
	for _, f := range funcs {
		tests := []struct {
			input, want []string
		}{
			{[]string{"a"}, []string{"a"}},
			{[]string{"a", "a"}, []string{"a"}},
			{[]string{"a", "b", "b", "a"}, []string{"a", "b", "a"}},
			{[]string{"a", "b", "a", "b"}, []string{"a", "b", "a", "b"}},
			{[]string{"a", "b", "b", "b", "a"}, []string{"a", "b", "a"}},
			{[]string{"a", "b", "b", "b", "a", "a", "b", "b", "b"}, []string{"a", "b", "a", "b"}},
			{[]string{"a", "a", "a", "a", "b", "b", "b", "b"}, []string{"a", "b"}},
		}

		for _, test := range tests {
			orig := make([]string, len(test.input))
			copy(orig, test.input)

			t.Run(f.name, func(t *testing.T) {
				if got := f.f(test.input); !eq(got, test.want) {
					t.Errorf("%s(%v) got %v, want %v", f.name, orig, got, test.want)
				}
			})

		}
	}

}

func eq(a, b []string) bool {
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

func BenchmarkDup(b *testing.B) {
	for _, f := range funcs {
		benchmarks := []struct {
			name    string
			strings []string
		}{
			{"LongSingleRun", []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a"}},
			{"RepeatingRuns", []string{"a", "b", "b", "c", "c", "c", "a", "b", "b", "c", "c", "c", "a", "b", "b", "c", "c", "c", "a", "b", "b", "c", "c", "c", "a", "b", "b", "c", "c", "c", "a", "b", "b", "c", "c", "c", "a", "b", "b", "c", "c", "c"}},
			{"NoRuns", []string{"a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c"}},
		}

		for _, benchmark := range benchmarks {
			b.Run(fmt.Sprintf("%s-%s", f.name, benchmark.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f.f(benchmark.strings)
				}
			})
		}
	}
}
