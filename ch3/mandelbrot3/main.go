// Mandelbrot3 measures the difference in performance and accuracy of rendering the
// mandelbrot set using various representations of numbers.
package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/masonelmore/gopl/ch3/mandelbrot3/mandelbrot"
	"github.com/masonelmore/gopl/ch3/mandelbrot3/mandelbrot/bigfloat"
	"github.com/masonelmore/gopl/ch3/mandelbrot3/mandelbrot/onetwoeight"
	"github.com/masonelmore/gopl/ch3/mandelbrot3/mandelbrot/sixtyfour"
)

func main() {
	const width, height = 1024, 1024
	const zoomPerIter = 0.1

	var renderFuncs = []struct {
		f       mandelbrot.Renderer
		name    string
		maxIter int
	}{
		{sixtyfour.Render, "complex64", 5},
		{onetwoeight.Render, "complex128", 14},
		{bigfloat.Render, "bigfloat", 14},
		// {bigrat.Render, "bigrat", 20},
	}

	for _, f := range renderFuncs {
		var xmin, ymin, xmax, ymax float64 = -2, -2, 2, 2
		currentzoom := 1.0
		for i := 0; i < f.maxIter; i++ {
			xmin, ymin, xmax, ymax = zoom(xmin, ymin, xmax, ymax, -1.5, 0, zoomPerIter)
			currentzoom *= zoomPerIter

			img := f.f(xmin, ymin, xmax, ymax, width, height)
			filename := fmt.Sprintf("images/%s-%.0e.png", f.name, currentzoom)
			err := saveImage(img, filename)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	}
}

func zoom(x0, y0, x1, y1, tox, toy, zoom float64) (xmin, ymin, xmax, ymax float64) {
	w := (x1 - x0) * zoom
	h := (y1 - y0) * zoom
	xmin = tox - w/2
	xmax = tox + w/2
	ymin = toy - h/2
	ymax = toy + h/2
	return xmin, ymin, xmax, ymax
}

func saveImage(img image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		return err
	}

	return nil
}
