package main

import "testing"

func TestDiff(t *testing.T) {
	var tests = []struct {
		a, b [32]byte
		want int
	}{
		{[32]byte{0: 1}, [32]byte{0: 1}, 0},
		{[32]byte{0: 1}, [32]byte{0: 2}, 2},
		{[32]byte{31: 7}, [32]byte{0: 5}, 5},
	}

	for _, test := range tests {
		if got := diff(&test.a, &test.b); got != test.want {
			t.Errorf("diff(%v, %v), got %d, want %d", test.a, test.b, got, test.want)
		}
	}
}
