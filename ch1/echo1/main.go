// Echo1 prints the name of the command that invoked it along with its
// command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	cmd := os.Args[0]
	args := strings.Join(os.Args[1:], " ")
	fmt.Println(cmd, args)
}
