package main

import (
	"fmt"

	"github.com/masonelmore/gopl/ch2/tempconv/tempconv"
)

func main() {
	c := tempconv.Celsius(30)
	f := tempconv.Fahrenheit(86)
	k := tempconv.Kelvin(303.15)

	fmt.Printf("%v = %v = %v\n", c, tempconv.CToF(c), tempconv.CToK(c))
	fmt.Printf("%v = %v = %v\n", f, tempconv.FToC(f), tempconv.FToK(f))
	fmt.Printf("%v = %v = %v\n", k, tempconv.KToC(k), tempconv.KToF(k))
}
