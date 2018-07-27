package main

import "testing"

func TestAnagram(t *testing.T) {
	var tests = []struct {
		a, b string
		want bool
	}{
		{"", "", false},
		{"a", "", false},
		{"", "b", false},
		{"a", "a", false},
		{"abc", "abc", false},
		{"abc", "acbacb", false},
		{"aaa", "aba", false},
		{"abc", "cba", true},
		{"ab", "ba", true},
	}

	for _, test := range tests {
		if got := anagram(test.a, test.b); got != test.want {
			t.Errorf("anagram(%q, %q), got %v, want %v", test.a, test.b, got, test.want)
		}
	}
}
