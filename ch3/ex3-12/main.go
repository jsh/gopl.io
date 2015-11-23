package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("len(os.Args) != 2")
		os.Exit(1)
	}
	first := os.Args[1]
	second := os.Args[2]
	if len(first) != len(second) {
		fmt.Printf("%s and %s are not anagrams\n", first, second)
		os.Exit(1)
	}

	for _, r := range first {
		var n1, n2 int
		for _, rr := range first {
			if rr == r {
				n1++
			}
		}
		for _, rr := range second {
			if rr == r {
				n2++
			}
		}
		if n2 != n1 {
			fmt.Printf("%s and %s are not anagrams\n", first, second)
			os.Exit(1)
		}
	}
	fmt.Printf("%s and %s are anagrams!\n", first, second)
}
