// Hash prints the SHA256, SHA385, or SHA512 hash of its input.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	hashType := flag.Int("type", 256, "SHA 256, 384, or 512")
	flag.Parse()

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	switch *hashType {
	case 256:
		fmt.Printf("%x\n", sha256.Sum256(input))
	case 384:
		fmt.Printf("%x\n", sha512.Sum384(input))
	case 512:
		fmt.Printf("%x\n", sha512.Sum512(input))
	default:
		fmt.Fprintf(os.Stderr, "%d is not a valid hash type.\n", *hashType)
		os.Exit(1)
	}
}
