// Package surface computes an SVG rendering of a 3-D surface function.
package surface

import (
	"fmt"
	"io"
	"math"
)

const angle = math.Pi / 6 // angle of x, y axes (=30°)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type Surface struct {
	width   int     // canvas width in pixels
	height  int     // canvas height in pixels
	color   string  // surface color
	cells   int     // number of grid cells
	xyrange float64 // axis ranges (-xyrange..+xyrange)
	xyscale float64 // pixels per x or y unit
	zscale  float64 // pixels per z unit
	// angle   float64 // angle of x, y axes (=30°)
}

func NewSurface(w, h int, c string) *Surface {
	cells := 100
	xyrange := 30.0
	xyscale := float64(w) / 2 / xyrange
	zscale := float64(h) * 0.4
	return &Surface{w, h, c, cells, xyrange, xyscale, zscale}
}

func (s *Surface) Generate(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", s.color, s.width, s.height)
	for i := 0; i < s.cells; i++ {
		for j := 0; j < s.cells; j++ {
			ax, ay := s.corner(i+1, j)
			bx, by := s.corner(i, j)
			cx, cy := s.corner(i, j+1)
			dx, dy := s.corner(i+1, j+1)
			fmt.Fprintf(out, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func (s *Surface) corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i, j).
	x := s.xyrange * (float64(i)/float64(s.cells) - 0.5)
	y := s.xyrange * (float64(j)/float64(s.cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(s.width)/2 + (x-y)*cos30*s.xyscale
	sy := float64(s.height)/2 + (x+y)*sin30*s.xyscale - z*s.zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
