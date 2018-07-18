package main

import (
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		ints  []int
		shift int
		want  []int
	}{
		// Positive Shift shift
		{[]int{1}, 0, []int{1}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2}, 0, []int{1, 2}},
		{[]int{1, 2}, 1, []int{2, 1}},
		{[]int{1, 2}, 2, []int{1, 2}},
		{[]int{1, 2}, 3, []int{2, 1}},
		{[]int{1, 2, 3, 4}, 0, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4}, 1, []int{4, 1, 2, 3}},
		{[]int{1, 2, 3, 4}, 2, []int{3, 4, 1, 2}},
		{[]int{1, 2, 3, 4}, 3, []int{2, 3, 4, 1}},
		{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4}, 5, []int{4, 1, 2, 3}},

		// Negative shift
		{[]int{1}, -1, []int{1}},
		{[]int{1, 2}, -1, []int{2, 1}},
		{[]int{1, 2}, -2, []int{1, 2}},
		{[]int{1, 2, 3, 4}, -1, []int{2, 3, 4, 1}},
		{[]int{1, 2, 3, 4}, -2, []int{3, 4, 1, 2}},
		{[]int{1, 2, 3, 4}, -3, []int{4, 1, 2, 3}},
		{[]int{1, 2, 3, 4}, -4, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4}, -5, []int{2, 3, 4, 1}},
	}

	for _, test := range tests {
		if got := rotate(test.ints, test.shift); !equals(got, test.want) {
			t.Errorf("rotate(%v, %v) got %v, want %v", test.ints, test.shift, got, test.want)
		}
	}
}

func TestRotateCopy(t *testing.T) {
	tests := []struct {
		ints  []int
		shift int
	}{
		// {[]int{}, 0},
		{[]int{1}, 0},
		{[]int{1, 2}, 0},
		{[]int{1, 2, 3}, 1},
	}

	for _, test := range tests {
		got := rotate(test.ints, test.shift)

		// Change got to check if the original has not been modified.
		got[0] = 9
		if equals(got, test.ints) {
			t.Errorf("rotate(%v, %v) did not return a new slice.", test.ints, test.shift)
		}
	}
}

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
