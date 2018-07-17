package main

import (
	"fmt"
)

const (
	KB = 1000
	MB = KB * KB
	GB = MB * MB
	TB = GB * GB
	PB = TB * TB
	EB = PB * PB
	ZB = EB * EB
	YB = ZB * ZB
)

func main() {
	fmt.Printf("KB: %v\nMB: %v\nGB: %v\nTB: %v\nPB: %v\nEB: %v\nEiB: %v\n", KB, MB, GB, TB, PB, EB, ZB, YB)
}
