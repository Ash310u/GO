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
	// Using Print line: Println
	if err != nil {
		// Both lines print the error, but there's a subtle difference:
		// fmt.Println(err) prints the error, which uses the Error() method internally, so it prints the error message.
		// fmt.Println(err.Error()) calls Error() explicitly and prints the resulting string.
		// In practice, for standard errors, these produce the same result.
		fmt.Println(err)
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Println("The result of the integer division is", result)
	} else {
		// Using Printf
		fmt.Printf("The result of the integer division is %v with remainder %v\n", result, remainder)
	}

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
