// Package popcount computes the number of set bits.
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountTable returns the population count (number of set bits) of x.
func PopCountTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountLoop returns the population count (number of set bit) of x.
func PopCountLoop(x uint64) int {
	var p int
	for i := 0; i < 8; i++ {
		p += int(pc[byte(x>>byte(i*8))])
	}
	return p
}

// PopCountShift returns the population count (number of set bits) of x.
func PopCountShift(x uint64) int {
	var p int
	for i := 0; i < 64; i++ {
		p += int(x >> byte(i) & 1)
	}
	return p
}

// PopCountClear returns the population count (number of set bits) of x.
func PopCountClear(x uint64) int {
	var p int
	for x > 0 {
		x &= (x - 1)
		p++
	}
	return p
}
