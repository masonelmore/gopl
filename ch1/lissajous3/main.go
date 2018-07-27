// Lissajous3 generates GIF animations of random lissajous figures and serves
// them through HTTP.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// default values for lissajous
	var (
		cycles  = 5   // number of complete x oscillator revolutions
		size    = 100 // image canvas covers [-size..+size]
		nframes = 64  // number of animation frames
		delay   = 8   //  delay between frames in 10ms units
	)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	// TODO: data validation for negative numbers and the like
	if c, err := strconv.Atoi(r.FormValue("cycles")); err == nil {
		cycles = c
	}
	if s, err := strconv.Atoi(r.FormValue("size")); err == nil {
		size = s
	}
	if n, err := strconv.Atoi(r.FormValue("nframes")); err == nil {
		nframes = n
	}
	if d, err := strconv.Atoi(r.FormValue("delay")); err == nil {
		delay = d
	}

	lissajous(w, cycles, size, nframes, delay)
}

func lissajous(out io.Writer, cycles, size, nframes, delay int) {
	const res = 0.001 // angular resolution

	freq := rand.Float64() * 3.0 // relative freqency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			px := size + int(x*float64(size)+0.5)
			py := size + int(y*float64(size)+0.5)
			img.SetColorIndex(px, py, blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
