// Unitconv converts its numeric arguments to various temperatures, lengths, and
// weights.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/masonelmore/gopl/ch2/tempconv/tempconv"
	"github.com/masonelmore/gopl/ch2/unitconv/lenconv"
	"github.com/masonelmore/gopl/ch2/unitconv/weightconv"
)

func main() {
	var args []string
	if len(os.Args) < 2 {
		fmt.Println("press ctrl-z or q and enter when you are done entering numbers.")
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("number: ")
		for scanner.Scan() {
			text := scanner.Text()
			if text == "q" {
				break
			}
			args = append(args, text)
			fmt.Print("number: ")
		}
	} else {
		args = os.Args[1:]
	}

	for _, arg := range args {
		n, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(n)
		c := tempconv.Celsius(n)
		ft := lenconv.Feet(n)
		m := lenconv.Meter(n)
		lb := weightconv.Pound(n)
		kg := weightconv.Kilogram(n)
		fmt.Printf("%s = %s, %s = %s\n%s = %s, %s = %s\n%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c),
			ft, lenconv.FToM(ft), m, lenconv.MToF(m),
			lb, weightconv.PToK(lb), kg, weightconv.KToP(kg))
	}
}
