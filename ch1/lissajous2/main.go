// Lissajous2 generates GIF animations of random lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = makePalette()

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 256   // number of animation frames
		delay   = 8     //  delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative freqency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		// img := image.NewRGBA(rect)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			px := size + int(x*size+0.5)
			py := size + int(y*size+0.5)
			colorIndex := uint8(i)
			img.SetColorIndex(px, py, colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func makePalette() []color.Color {
	var p []color.Color
	step := 0xffffff / 256
	for i := 0; i < 256; i++ {
		c := i * step
		b := uint8(c)
		g := uint8(c >> 8)
		r := uint8(c >> 16)
		p = append(p, color.RGBA{r, g, b, 255})
	}
	return p
}
