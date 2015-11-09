// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	names := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines("/dev/stdin", counts, names)
	} else {
		for _, arg := range files {
			countLines(arg, counts, names)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, names[line])
		}
	}
}

func countLines(filename string, counts map[string]int, names map[string]string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		return
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		names[input.Text()] += " " + filename
	}
	f.Close()
	// NOTE: ignoring potential errors from input.Err()
}

//!-
