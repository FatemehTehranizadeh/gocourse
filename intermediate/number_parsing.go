package main

import (
	"fmt"
	"strconv"
)


func main(){
	/* Number Parsing: Process of converting textual representations of numbers in to
	their corresponding numeric values.
	It's done using STR.CONV package.

	*/

	// Converting string to integer
	numStr := "123456"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error in converting string to integer: ", err)
	}
	fmt.Printf("The converted string is %d and its type is %T\n", num, num)

	parsedInt, err := strconv.ParseInt(numStr, 16, 32) // It returns the number in base 16
	if err != nil {
		fmt.Println("Error in converting string to int32: ", err)
	}
	fmt.Printf("The converted string is %d and its type is %T\n", parsedInt, parsedInt)

	floatStr := "2.46677"
	floatNum, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Error in converting string to float64: \n ", err)
	}
	fmt.Printf("The parsed float is %.6f and its type is %T\n", floatNum, floatNum)

	binaryNum := "-10101"
	decimalNum, err := strconv.ParseInt (binaryNum, 2, 32)
	if err != nil {
		fmt.Println("Error in converting string to int32: ", err)
	}
	fmt.Printf("The converted string is %d and its type is %T\n", decimalNum, decimalNum)




}