// Package weightconv performs Pound and Kilogram weight computations.
package weightconv

import (
	"fmt"
)

type Pound float64
type Kilogram float64

// PToK converts a weight in Pounds to Kilograms.
func PToK(p Pound) Kilogram { return Kilogram(p * 0.45359237) }

// KToP converts a weight in Kilograms to Pounds.
func KToP(k Kilogram) Pound { return Pound(k / 0.45359237) }

func (p Pound) String() string    { return fmt.Sprintf("%glbs", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
