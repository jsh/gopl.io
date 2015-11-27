package main

// Print the SHA256 of the stdin by default,
// but accept command-line flag to print SHA384 or SHA512 instead.

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	sha := 256
	if len(os.Args) == 2 {
		sha, _ = strconv.Atoi(os.Args[1])
	}
	b, err := ioutil.ReadFile("/dev/stdin")
	if err != nil {
		fmt.Printf("can't read stdin")
		os.Exit(1)
	}
	switch sha {
	case 256:
		c := sha256.Sum256(b)
		fmt.Printf("The SHA%d hash is %x\n", sha, c)
	case 384:
		c := sha512.Sum384(b)
		fmt.Printf("The SHA%d hash is %x\n", sha, c)
	case 512:
		c := sha512.Sum512(b)
		fmt.Printf("The SHA%d hash is %x\n", sha, c)
	default:
		fmt.Printf("Unknown SHA hash SHA%d\n", sha)
		os.Exit(1)
	}
}
