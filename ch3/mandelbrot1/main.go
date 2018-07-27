// Mandelbrot1 emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return kolor(contrast * n)
		}
	}
	return color.Black
}

func kolor(n uint8) color.Color {
	// http://www.andrewnoske.com/wiki/Code_-_heatmaps_and_color_gradients
	var r0, g0, b0 uint8 = 255, 255, 255
	var r1, g1, b1 uint8 = 255, 69, 0
	var r2, g2, b2 uint8 = 0, 0, 0

	r := (r2-r1-r0)*n + r0
	g := (g2-g1-g0)*n + g0
	b := (b2-b1-b0)*n + b0
	return color.RGBA{r, g, b, 255}
}
