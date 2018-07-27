// Hashdiff counts the number of bits that are different in two SHA256 hashes.
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	n := diff(&c1, &c2)
	fmt.Println(n)
}

func diff(a *[32]byte, b *[32]byte) int {
	var n int
	for i, v := range a {
		n += PopCount(v ^ b[i])
	}
	return n
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x byte) int {
	var p int
	for x > 0 {
		x = x & (x - 1)
		p++
	}
	return p
}
