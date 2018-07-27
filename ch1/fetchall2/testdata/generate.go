// testdata prints a number of urls from the alexa top 1m csv.  The csv can be
// found at http://s3.amazonaws.com/alexa-static/top-1m.csv.zip
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: generate <filename> <count>")
		os.Exit(1)
	}

	filename := os.Args[1]
	count, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: invalid value for count: %s\n", os.Args[2])
		os.Exit(1)
	}
	if count < 1 {
		fmt.Fprintf(os.Stderr, "error: invalid value for count: %s\n", os.Args[2])
		os.Exit(1)
	}

	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)
	for i := 0; i < count; i++ {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		url := "http://" + strings.Split(line, ",")[1]
		fmt.Println(url)
	}
}
