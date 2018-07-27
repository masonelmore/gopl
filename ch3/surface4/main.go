// Surface4 renders a 3-D surface using SVG over HTTP.
package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/masonelmore/gopl/ch3/surface4/surface"
)

const (
	defaultWidth  = 600
	defaultHeight = 320
	defaultColor  = "grey"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	width, err := strconv.Atoi(r.FormValue("width"))
	if err != nil {
		width = defaultWidth
	}
	height, err := strconv.Atoi(r.FormValue("height"))
	if err != nil {
		height = defaultHeight
	}
	color := r.FormValue("color")
	if color == "" {
		// NOTE: Only color names (blue, yellow, brown, etc.) work because using
		// a # to specify a hex color messes with the request.  I could possibly
		// check if it's valid hex and toss a # in front for a quick fix or just
		// use a urlencoded # (e.g. %2300FF00).
		color = defaultColor
	}
	s := surface.NewSurface(width, height, color)
	s.Generate(w)
}
