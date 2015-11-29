// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	fmt.Println(rotate(s, 3)) // "[2 3 4 5 0 1]"

}

// rotate rotates a slice of ints.
func rotate(s []int, n int) []int {
	t := make([]int, len(s))
	// Rotate s left by two positions.
	copy(t, s[n:])
	copy(t[len(s)-n:], s[:n])
	return t
}
