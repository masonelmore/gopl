// ex3.7 emits a PNG of a simple fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -1, -1, 1, 1
	width, height          = 1024, 1024
)

var (
	roots = []complex128{
		1 + 0i,
		0 + 1i,
		-1 + 0i,
		0 + -1i,
	}
	colors = []color.Color{
		color.RGBA{255, 8, 0, 255},
		color.RGBA{84, 255, 13, 255},
		color.RGBA{2, 255, 237, 255},
		color.RGBA{85, 36, 232, 255},
	}
	eps = math.Nextafter(1, 2) - 1
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			c := newton(z)
			if c == nil {
				continue
			}
			img.Set(px, py, c)
		}
	}
	png.Encode(os.Stdout, img)
}

func newton(z complex128) color.Color {
	const iterations = 12

	for i := 0; i < iterations; i++ {
		for r := 0; r < len(roots); r++ {
			d := z - roots[r]
			if cmplx.Abs(d) < eps {
				intensity := 1 - (float64(i) / (iterations))
				col := shade(intensity, colors[r])
				return col
			}
			z -= (z - 1/(z*z*z)) / 4
		}
	}

	return color.Black
}

func shade(val float64, colr color.Color) color.Color {
	r, g, b, a := colr.RGBA()
	rr := float64(r) * val
	gg := float64(g) * val
	bb := float64(b) * val
	if val > 1 {
		fmt.Fprintln(os.Stderr, val)
	}
	return color.RGBA{uint8(rr / 0x101), uint8(gg / 0x101), uint8(bb / 0x101), uint8(a)}
}
