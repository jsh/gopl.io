package main

import "fmt"

const hello = "Hello, 世界"

// const hello = "hello, world"

func main() {
	b := []byte(hello)
	reverse(b)
	fmt.Printf("%v %v\n", []byte(hello), b)
}

func reverse(b []byte) {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
}
