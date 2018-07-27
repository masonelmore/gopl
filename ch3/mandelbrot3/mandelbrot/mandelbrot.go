// Package mandelbrot provides implementations for rendering the mandelbrot set
// using following number representations: complex64, complex128, big.Float, or
// big.Rat.
package mandelbrot

import (
	"image"
)

// A Renderer produces an image of the mandelbrot set.
type Renderer func(float64, float64, float64, float64, int, int) image.Image
