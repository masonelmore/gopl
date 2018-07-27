package sixtyfour

import (
	"image"
	"image/color"
	"math/cmplx"
)

// Render generates an image of the mandelbrot set.
func Render(xmin, ymin, xmax, ymax float64, width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(float32(x), float32(y))
			img.Set(px, py, mandelbrot(z))
		}
	}
	return img
}

func mandelbrot(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		// TODO: Does converting to complex128 fudge the results in any way?
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
