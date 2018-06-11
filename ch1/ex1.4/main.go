// ex1.4 prints the count and text of lines that appear more than once in the
// input files, followed by their filenames.  It reads from stdin or from a list
// of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type count struct {
	count int
	files []string
}

func main() {
	counts := make(map[string]count)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, c := range counts {
		if c.count > 1 {
			fmt.Printf("%d\t%s\t%s\n", c.count, line, strings.Join(c.files, ", "))
		}
	}
}

func countLines(f *os.File, counts map[string]count) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		c := counts[input.Text()]
		c.count++

		var found bool
		for _, fn := range c.files {
			if fn == f.Name() {
				found = true
				break
			}
		}
		if !found {
			c.files = append(c.files, f.Name())
		}

		counts[input.Text()] = c
	}
	// NOTE: ignoring potential errors from input.Err()
}
