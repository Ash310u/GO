package main

import "fmt"

func main() {
	var numerator, denominator float32

	PrintMe("What do you want to divide:")
	fmt.Scan(&numerator, &denominator)

	result := intDivision(numerator, denominator)
	fmt.Println("Result:", result)
}

func PrintMe(printValue string) {
	fmt.Println(printValue)
}

// Yes, the 'int' before the curly bracket specifies the return type of the function.
// Here is the function demonstrating the return type:
func intDivision(numerator float32, denominator float32) float32 { // 'int' indicates the function returns an int
	return numerator / denominator
}
