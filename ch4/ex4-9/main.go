// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 96.

// Dedup prints only one instance of each line; duplicates are removed.
package main

import (
	"bufio"
	"fmt"
	"os"
)

//!+
func main() {
	words := make(map[string]int) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		words[word]++
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}

	for word, n := range words {
		fmt.Printf("%s: %d\n", word, n)
	}
}

//!-
