// Fetchall2 fetches URLs in parallel and reports their times and sizes.
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("results", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
	if err != nil {
		fmt.Fprintf(os.Stderr, "while opening file: %v\n", err)
		os.Exit(1)
	}
	fw := bufio.NewWriter(f)

	start := time.Now()
	fmt.Fprintln(fw, start.String())
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	for range os.Args[1:] {
		fmt.Fprintln(fw, <-ch) // receive from channel ch
	}
	fmt.Fprintf(fw, "%.2fs elapsed\n\n", time.Since(start).Seconds())
	fw.Flush()
	f.Close()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		ch <- fmt.Sprintf("%.2fs    ERROR  %s", time.Since(start).Seconds(), url)
		return
	}

	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	ch <- fmt.Sprintf("%.2fs  %7d  %s", time.Since(start).Seconds(), nBytes, url)
}
