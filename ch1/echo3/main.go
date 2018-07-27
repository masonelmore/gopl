// Echo3 measures the difference in running time between using string
// concatination and strings.Join for our echo program.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]

	start := time.Now()
	echoConcat(args)
	fmt.Printf("%dns elapsed for echoConcat\n", time.Since(start).Nanoseconds())

	start = time.Now()
	echoJoin(args)
	fmt.Printf("%dns elapsed for echoJoin\n", time.Since(start).Nanoseconds())
}

func echoConcat(args []string) string {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

func echoJoin(args []string) string {
	return strings.Join(args, " ")
}
