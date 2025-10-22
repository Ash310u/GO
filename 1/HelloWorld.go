package main

import "fmt"

func main() {
	// single-line comments start with "//"
	// comments are just for documentation - they don't execute
	fmt.Println("hello world")

	// To append data in Go, use the built-in append function, usually with slices.
	s := []string{"hello"}
	s = append(s, "world")
	fmt.Println(s) // Output: [hello world]
}
