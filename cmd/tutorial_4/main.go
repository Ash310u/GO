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

	// 1. Literal initialization (most common)
	slice1 := []string{"apple", "banana", "cherry"}
	// 2. Using make() to create a slice with length (and optional capacity)
	slice2 := make([]int, 3) // length 3, capacity 3, zero-initialized
	slice2[0] = 11
	slice2[1] = 22
	slice2[2] = 33
	// 3. Creating a slice from an array
	myArr := [4]float64{1.0, 2.2, 3.3, 4.4}
	slice3 := myArr[1:3] // slice containing elements at index 1 and 2 (2.2, 3.3)
	// 4. Declaring a nil slice
	var slice4 []int
	// 5. Creating an empty slice using a literal
	slice5 := []byte{}

	// Show different ways to create slices
	fmt.Println("Slice1 (literal):", slice1)
	fmt.Println("Slice2 (using make):", slice2)
	fmt.Println("Slice3 (from array):", slice3)
	fmt.Println("Slice4 (nil):", slice4)
	fmt.Println("Slice5 (empty):", slice5)

	// Append to a slice
	slice1 = append(slice1, "date")
	fmt.Println("Slice1 after append:", slice1)
	// slice1[1:4] creates a new slice from index 1 up to 4 (exclusive)
	fmt.Println("Slice1[1:4]:", slice1[1:4])

	// ---------------- Maps ----------------
	// The 'make' function is used in Go to create slices, maps, and channels.
	// Here, it creates an empty map with keys of type string and values of type int.
	m := make(map[string]int)
	m["foo"] = 42
	m["bar"] = 27
	fmt.Println("Map:", m)
	val, exists := m["foo"]
	if exists {
		fmt.Println("Value for key 'foo':", val)
	}

	// Looping
	fmt.Println("\nLoop: array values")
	for i := 0; i < len(arr1); i++ {
		fmt.Println(i, arr1[i])
	}

	fmt.Println("\nLoop: slice1 values with range")
	for index, val := range slice1 {
		fmt.Printf("Index %d: %s\n", index, val)
	}

	fmt.Println("\nLoop: map key-values")
	for k, v := range m {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

	// Loop control
	fmt.Println("\nLoop control with continue/break")
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("skip 2")
			continue
		}
		if i == 4 {
			fmt.Println("break at 4")
			break
		}
		fmt.Println("i:", i)
	}

	// Struct
	// In Go, struct field names must start with a capital letter if you want to access them from other packages (exported).
	// If the struct is only used within the same package, you may use lowercase, but then the fields will not be accessible from outside.
	type Person struct {
		Name string
		Age  int
		// name string // unexported (accessible only within this package)
		// age  int    // unexported (accessible only within this package)
	}
	person := Person{Name: "Alice", Age: 30}
	fmt.Println("\nStruct value:", person)
	fmt.Println("Person's Name:", person.Name)
	fmt.Println("Person's Age:", person.Age)

	// Slice of structs
	people := []Person{
		{Name: "Bob", Age: 25},
		{Name: "Carol", Age: 27},
	}

	people = append(people, person)
	fmt.Println("\nSlice of structs:")
	for i, p := range people {
		fmt.Printf("Name %d : %s, Age: %d\n", i, p.Name, p.Age)
	}
}
