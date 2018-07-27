package squash

import (
	"testing"
)

func TestSquash(t *testing.T) {
	// TODO: This seems like an excessive number of tests.  I should spend more
	// time thinking about what exactly I need to test.
	tests := []struct {
		input, want string
	}{
		{"a", "a"},
		{"aa", "aa"},
		{"", ""},
		{" ", " "},
		{"  ", " "},
		{"   ", " "},
		{"    ", " "},
		{"a ", "a "},
		{"a  ", "a "},
		{"a   ", "a "},
		{"a    ", "a "},
		{" a ", " a "},
		{"  a  ", " a "},
		{"   a   ", " a "},
		{"    a    ", " a "},
		{"a b c d", "a b c d"},
		{"a  b    c     d", "a b c d"},
		{"a  b    c     d     ", "a b c d "},
		{"\n", "\n"},
		{"\n\n", " "},
		{"\t\n\v\f\r", " "},
		{"\n\n\n\na\n", " a\n"},
		{"\n\n\n\na\n\n\n\n", " a "},
		{"a\na", "a\na"},
		{"a\n\na", "a a"},
		{"a\n\na\n", "a a\n"},
		{"a\n\na\n\n", "a a "},
		{"\u0085", "\u0085"},
		{"\u0085\u0085", " "},
		{"\u00A0", "\u00A0"},
		{"\u00A0\u00A0", " "},
		{"\u00A0\u00A0\u0085\u0085", " "},
		{"😨", "😨"},
		{"😨 ", "😨 "},
		{"😨\n", "😨\n"},
		{"😨\u00A0", "😨\u00A0"},
		{"😨\u00A0\u00A0", "😨 "},
		{"😨\u00A0\u00A0😨", "😨 😨"},
		{"𠜎", "𠜎"},
		{"𠜎 ", "𠜎 "},
		{" 𠜎 ", " 𠜎 "},
		{" 𠜎 ", " 𠜎 "},
		{" 𠜎 ", " 𠜎 "},
		{" \u00A0\u00A0𠜎\u00A0\u00A0 ", " 𠜎 "},
		{" \u00A0\u00A0𠜎\u00A0\u00A0 𠜎", " 𠜎 𠜎"},
	}

	for _, test := range tests {
		t.Run("squash", func(t *testing.T) {
			input := []byte(test.input)
			want := []byte(test.want)
			if got := squash(input); !eq(got, want) {
				t.Errorf("squash(%q), got %q, want %q", test.input, got, want)
			}
		})
	}
}

func BenchmarkDup(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"LongSingleRun", "                                                                                                    "},
		{"RepeatingRuns", "a  b   c    a  b   c    a  b   c    a  b   c    a  b   c    a  b   c    a  b   c    a  b   c    a  b"},
		{"NoRuns", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
	}

	for _, benchmark := range benchmarks {
		bytes := []byte(benchmark.input)
		b.Run("squash", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				squash(bytes)
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
