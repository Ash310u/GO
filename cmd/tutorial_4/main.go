package main

import (
	"fmt"
)

func main() {
	// ---------------- Arrays ----------------
	// 1. Explicit type and length with initialization
	var arr1 [5]int32 = [5]int32{1, 2, 3, 4, 5}
	// 2. Type inference, length specified
	var arr2 = [5]int32{10, 20, 30, 40, 50}
	// 3. Shorthand using :=
	arr3 := [5]int{100, 200, 300, 400, 500}
	// 4. Compiler infers the length
	// The '...' in [...]string tells Go to infer the array length based on the number of elements provided.
	arr4 := [...]string{"Go", "Rust", "Python"} // Here, arr4 becomes an array of length 3, initialized with the given values.
	// 5. Declaring without initializing, zero values
	var arr5 [3]float64
	arr5[0] = 1.1
	arr5[1] = 1.2
	arr5[2] = 1.3
	fmt.Println("Array1:", arr1)
	fmt.Println("Array2:", arr2)
	fmt.Println("Array3:", arr3)
	fmt.Println("Array4:", arr4)
	fmt.Println("Array5:", arr5)

	// ---------------- Slices ----------------
	slice := []string{"apple", "banana", "cherry"}
	slice = append(slice, "date")
	fmt.Println("Slice after append:", slice)
	// How does this work?
	// slice[1:3] creates a new slice starting from index 1 (inclusive) up to index 4 (exclusive).
	// In our example, slice = []string{"apple", "banana", "cherry", "date"}
	// slice[1:3] will be []string{"banana", "cherry", "date"}
	fmt.Println("Slice[1:3]:", slice[1:4])

	// Maps
	m := make(map[string]int)
	m["foo"] = 42
	m["bar"] = 27
	fmt.Println("Map:", m)
	val, exists := m["foo"]
	if exists {
		fmt.Println("Value for key 'foo':", val)
	}

}
