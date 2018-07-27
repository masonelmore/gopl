package bigfloat

import (
	"image"
	"image/color"
	"math/big"
)

// Render generates an image of the mandelbrot set.
func Render(xmin, ymin, xmax, ymax float64, width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var cr, ci big.Float
	for py := 0; py < height; py++ {
		y := ((float64(py) / float64(height)) * (ymax - ymin)) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			cr.SetFloat64(x)
			ci.SetFloat64(y)
			img.Set(px, py, mandelbrot(&cr, &ci))
		}
	}
	return img
}

func mandelbrot(cr, ci *big.Float) color.Color {
	// https://randomascii.wordpress.com/2011/08/13/faster-fractals-through-algebra/
	const iterations = 200
	const contrast = 15

	var zr, zi, zrzi, zrsqr, zisqr big.Float
	for n := uint8(0); n < iterations; n++ {
		zrzi.Add(&zr, &zi)
		zi.Mul(&zrzi, &zrzi).Sub(&zi, &zrsqr).Sub(&zi, &zisqr).Add(&zi, ci)
		zr.Sub(&zrsqr, &zisqr).Add(&zr, cr)
		zrsqr.Mul(&zr, &zr)
		zisqr.Mul(&zi, &zi)
		if new(big.Float).Add(&zrsqr, &zisqr).Cmp(big.NewFloat(4.0)) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
