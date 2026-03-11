package main

import (
	"fmt"
)

const (
	KB = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB // Too big to print!
	YB = ZB * KB // Too big to print!
)

func main() {
	fmt.Printf("KB: %v\nMB: %v\nGB: %v\nTB: %v\nPB: %v\nEB: %v\n", KB, MB, GB, TB, PB, EB)
}
