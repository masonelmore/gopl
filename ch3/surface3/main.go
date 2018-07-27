// Surface3 computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _, ok := corner(i+1, j)
			if !ok {
				continue
			}
			bx, by, _, ok := corner(i, j)
			if !ok {
				continue
			}
			cx, cy, _, ok := corner(i, j+1)
			if !ok {
				continue
			}
			dx, dy, color, ok := corner(i+1, j+1)
			if !ok {
				continue
			}
			fmt.Printf("<polygon style='stroke: %s' points='%g,%g,%g,%g,%g,%g,%g,%g' />\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Printf("</svg>")
}

func corner(i, j int) (float64, float64, string, bool) {
	// Find point (x,y) at corner of cell (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsNaN(z) {
		return 0, 0, "", false
	}
	color := color(-0.2, 1.0, z)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, color, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func color(min, max, z float64) string {
	// https://stackoverflow.com/a/20792531
	ratio := 2 * (z - min) / (max - min)
	b := uint8(math.Max(0, 255*(1-ratio)))
	r := uint8(math.Max(0, 255*(ratio-1)))
	g := uint8(255 - b - r)

	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}
