// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
//!+
package main

import (
	"crypto/sha256"
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	sum := 0
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	for i := 0; i < len(c1); i++ {
		sum += PopCount(uint64(c1[i] ^ c2[i]))
	}
	fmt.Printf("There are %d bits different\n", sum)
}

//!-
