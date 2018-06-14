// Package lenconv performs Feet and Meter length computations.
package lenconv

import (
	"fmt"
)

type Feet float64
type Meter float64

// FToM converts a length in Feet to Meters.
func FToM(f Feet) Meter { return Meter(f / 3.2808) }

// MToF converts a length in Meters to Feet.
func MToF(m Meter) Feet { return Feet(m * 3.2808) }

func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
