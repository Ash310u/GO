package main

import "fmt"

func main() {
	var numerator, denominator float32

	PrintMe("What do you want to divide:")
	fmt.Scan(&numerator, &denominator)

	result, quotient := intDivisionAndQuotient(numerator, denominator)
	// Using Print line: Println
	fmt.Println("Result:", result)
	fmt.Println("Quotient:", quotient)

	// Using Printf
	fmt.Printf("The result of the integer division is %v with remainder %v\n", result, quotient)
}

func PrintMe(printValue string) {
	fmt.Println(printValue)
}

// Yes, the 'int' before the curly bracket specifies the return type of the function.
// Here is the function demonstrating the return type:
func intDivision(numerator float32, denominator float32) float32 { // 'int' indicates the function returns an int
	return numerator / denominator
}

func intDivisionAndQuotient(numerator float32, denominator float32) (float32, int) {
	divident := numerator / denominator
	quotient := int(numerator) % int(denominator)
	return divident, quotient
}
