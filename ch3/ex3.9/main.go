// ex3.9 renders fractals over HTTP.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	Xmin, Ymin, Xmax, Ymax = -2, -2, 2, 2
	Width, Height          = 1024, 1024
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// NOTE: My error handling isn't pretty, but I think it's fine for this problem.
	x, err := strconv.ParseFloat(r.FormValue("x"), 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	y, err := strconv.ParseFloat(r.FormValue("y"), 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	scale, err := strconv.ParseFloat(r.FormValue("scale"), 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	xmin, ymin, xmax, ymax := zoom(x, y, scale)
	img := render(xmin, ymin, xmax, ymax, Width, Height)
	png.Encode(w, img)
}

func zoom(x, y, scale float64) (xmin, ymin, xmax, ymax float64) {
	w := (Xmax - Xmin) * scale
	h := (Ymax - Ymin) * scale
	xmin = x - w/2
	xmax = x + w/2
	ymin = y - h/2
	ymax = y + h/2
	return xmin, ymin, xmax, ymax
}

func render(xmin, ymin, xmax, ymax float64, width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
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
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
