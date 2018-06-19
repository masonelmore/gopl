// Mandelbrot emits a PNG image of the Mandelbrot fractal with grid based
// supersampling.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const xmin, ymin, xmax, ymax = -2, -2, 2, 2

func main() {
	fSuper, err := os.OpenFile("images/supersample.png", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer fSuper.Close()
	img := supersample(1024, 1024)
	err = png.Encode(fSuper, img)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fNorm, err := os.OpenFile("images/normal.png", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer fNorm.Close()
	img = normal(1024, 1024)
	err = png.Encode(fNorm, img)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func supersample(width, height int) image.Image {
	bigImg := normal(width*2, height*2)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		bpy := py * 2
		for px := 0; px < width; px++ {
			bpx := px * 2
			c0 := bigImg.At(bpx, bpy)
			c1 := bigImg.At(bpx+1, bpy)
			c2 := bigImg.At(bpx, bpy+1)
			c3 := bigImg.At(bpx+1, bpy+1)
			avg := avgColor(c0, c1, c2, c3)
			img.Set(px, py, avg)
		}
	}
	return img
}

func normal(width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	return img
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func avgColor(colors ...color.Color) color.Color {
	// NOTE:  I was getting some wonky colors when I first tried this.  It was
	// because the color values from color.RGBA() are converted to a 16-bit
	// value inside a uint32.  Dividing by 0x101 before setting the color value
	// fixes that problem.
	// http://jimsaunders.net/2015/05/22/manipulating-colors-in-go.html

	// On another NOTE:  Computing the average color is not as simple as I am
	// doing it here.  Here's another way to do it:
	// https://medium.com/@kevinsimper/how-to-average-rgb-colors-together-6cd3ef1ff1e5
	var rsum, gsum, bsum uint32
	for _, c := range colors {
		r, g, b, _ := c.RGBA()
		rsum += r
		gsum += g
		bsum += b
	}
	n := uint32(len(colors))
	r := rsum / n
	g := gsum / n
	b := bsum / n
	return color.RGBA{uint8(r / 0x101), uint8(g / 0x101), uint8(b / 0x101), 255}
}
