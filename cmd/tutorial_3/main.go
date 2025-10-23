package main

import (
	"errors"
	"fmt"
)

func main() {
	var numerator, denominator float32

	PrintMe("What do you want to divide:")
	fmt.Scan(&numerator, &denominator)

	result, remainder, err := intDivisionAndQuotient(numerator, denominator)

	if err != nil {
		// Both lines print the error, but there's a subtle difference:
		// fmt.Println(err) prints the error, which uses the Error() method internally, so it prints the error message.
		// fmt.Println(err.Error()) calls Error() explicitly and prints the resulting string.
		// In practice, for standard errors, these produce the same result.

		fmt.Println(err)         // Using Print line: Println
		fmt.Println(err.Error()) // Using Print line: Println
	} else if denominator == numerator || remainder == 0 {
		fmt.Println("The result of the integer division is", result)
	} else {
		// Using Printf
		fmt.Printf("The result of the integer division is %v with remainder %v\n", result, remainder)
	}

	// Example of using switch statement to handle various division scenarios
	switch {
	case err != nil:
		fmt.Println("Error occurred during division, can't show results in switch statement.")
	case denominator == 0:
		fmt.Println("Cannot divide by zero.")
	case denominator == numerator:
		fmt.Println("Numerator is equal to Denominator in switch: result =", result)
	case remainder == 0:
		fmt.Println("Perfect division (no remainder) in switch, result:", result)
	default:
		fmt.Printf("Switch: %v divided by %v gives quotient %v and remainder %v\n", numerator, denominator, result, remainder)
	}
	// Example: Switch on a string representing an operation
	// operation := "add"
	// switch operation {
	// case "add":
	// 	fmt.Println("Adding two numbers!")
	// case "subtract":
	// 	fmt.Println("Subtracting two numbers!")
	// case "multiply":
	// 	fmt.Println("Multiplying two numbers!")
	// case "divide":
	// 	fmt.Println("Dividing two numbers!")
	// default:
	// 	fmt.Println("Unknown operation!")
	// }

}

func PrintMe(printValue string) {
	fmt.Println(printValue)
}

// Yes, the 'int' before the curly bracket specifies the return type of the function.
// Here is the function demonstrating the return type:
// func intDivision(numerator float32, denominator float32) float32 { // 'int' indicates the function returns an int
// 	return numerator / denominator
// }

func intDivisionAndQuotient(numerator float32, denominator float32) (float32, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}
	divident := numerator / denominator
	remainder := int(numerator) % int(denominator)
	return divident, remainder, err
}
