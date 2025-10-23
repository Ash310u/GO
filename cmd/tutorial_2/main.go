// Default values in Go:
// Integers (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64): 0
// Floats (float32, float64): 0
// Rune (alias for int32): 0
// Boolean (bool): false
// String: ""
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Integer examples
	var int16Max int16 = 32766
	var int32Max int32 = 2147483647
	var int64Max int64 = 9223372036854775807
	fmt.Println("int16 max value:", int16Max) // output: int16 max value: 32766
	fmt.Println("int32 max value:", int32Max) // output: int32 max value: 2147483647
	fmt.Println("int64 max value:", int64Max) // output: int64 max value: 9223372036854775807

	// Float examples
	var float32Max float32 = 3.4028235e+38
	var float64Max float64 = 1.7976931348623157e+308
	fmt.Println("float32 max value:", float32Max) // output: float32 max value: 3.4028235e+38
	fmt.Println("float64 max value:", float64Max) // output: float64 max value: 1.7976931348623157e+308

	var int32Num int32 = 2
	var float32Num float32 = 2.1
	// Can't perform operation with mixtypes
	// var result float32 = float32 + int32
	// By Casting the number we can perform operation like that
	var result float32 = float32Num + float32(int32Num)
	fmt.Println("Result by casting types:", result) // output: Result by casting types: 4.1

	// Integer division (truncation) example
	var dividend int = 3
	var divisor int = 2
	var quotient int = dividend / divisor             // result is truncated toward zero
	fmt.Println("Integer division 3 / 2 =", quotient) // output: Integer division 3 / 2 = 1

	// Remainder (modulus) example
	var remainder int = dividend % divisor
	fmt.Println("Remainder of 3 % 2 =", remainder) // output: Remainder of 3 % 2 = 1

	// Strings
	// len returns the number of bytes, not the actual character count (for non-ASCII strings)
	fmt.Println("Length of 'A'(Letter) inside the vanilla ascii characters (in bytes):", len("A")) // output: Length of 'A'(Letter) inside the vanilla ascii characters (in bytes): 1
	fmt.Println("Length of 'γ'(Gamma) outside the vanilla ascii characters (in bytes):", len("γ")) // output: Length of 'γ'(Gamma) outside the vanilla ascii characters (in bytes): 2

	var text1 string = "hello"
	var text2 string = "world"
	var text3 string = text1 + " γ " + text2
	fmt.Println("text3: ", text3) // output: text3:  hello γ world

	fmt.Println("Length of text3 with len(): ", len(text3))                                       // output: Length of text3 with len():  13
	fmt.Println("Length of text3 with utf8.RuneCountInString(): ", utf8.RuneCountInString(text3)) // output: Length of text3 with utf8.RuneCountInString():  11

	// Rune
	var myRune1 rune = 'A'
	var myRune2 rune = 'a'
	fmt.Println(myRune1) // output: 65
	fmt.Println(myRune2) // output: 97

	// Boolean
	var myBoolean1 bool = true
	var myBoolean0 bool = false
	fmt.Println(myBoolean1) // output: true
	fmt.Println(myBoolean0) // output: false

	// Short Hand any variable declaration
	myVar := "ashu"
	fmt.Println(myVar)

	var1, var2 := "A", 1
	fmt.Println(var1, var2)

	const constVar string = "Const Variable"
	fmt.Println(constVar)
}
