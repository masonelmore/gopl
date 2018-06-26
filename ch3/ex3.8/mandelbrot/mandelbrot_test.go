package mandelbrot

import (
	"testing"

	"github.com/masonelmore/gopl/ch3/ex3.8/mandelbrot/bigfloat"
	"github.com/masonelmore/gopl/ch3/ex3.8/mandelbrot/onetwoeight"
	"github.com/masonelmore/gopl/ch3/ex3.8/mandelbrot/sixtyfour"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
)

func BenchmarkRenderFuncs(b *testing.B) {
	var funcs = []struct {
		f    Renderer
		name string
	}{
		{sixtyfour.Render, "complex64"},
		{onetwoeight.Render, "complex128"},
		{bigfloat.Render, "bigfloat"},
		// {bigrat.Render, "bigrat"},
	}

	var benchmarks = []struct {
		width, height int
		name          string
	}{
		{256, 256, "256"},
		{512, 512, "512"},
		{1024, 1024, "1024"},
		{2048, 2048, "2048"},
	}

	for _, f := range funcs {
		for _, bm := range benchmarks {
			b.Run(f.name+"-"+bm.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f.f(xmin, ymin, xmax, ymax, bm.width, bm.height)
				}
			})
		}
	}
}
