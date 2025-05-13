package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float32 = 10, 3
	var result float32

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Devision:", result)

	// result = a % b
	// fmt.Println("Remainder:", result)

	const p float64 = 22 / 7.0
	fmt.Println(p)

	//Overflow for signed and unsigned numbers
	var maxInt int64 = math.MaxInt64
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt) //Results a negative number

	
	var uMaxInt uint64 = math.MaxUint64
	
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt) //Results 0

	//Underflow for signed and unsigned numbers
	var smallFloat float64 = 1.0e-323
	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat) //Results 0



}