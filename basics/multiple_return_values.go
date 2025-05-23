package main

import (
	"errors"
	"fmt"
)


func main(){
	/* func functionName (parameter1 type1, parameter2, type2, ... ) (returnType1, returnType2, ...){
	code block
	return returnValue1, returnValue2, ...
	}
	*/
	q, r := divide(10,3)
	fmt.Println("Quitient:", q, "Reminder:",r)

	//The biggest benefit of multiple return values is error handling
	result, err := compare(-5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	 





}

func divide (a, b int) (quotient int, reminder int) {
	quotient = a / b
	reminder = a % b
	return quotient, reminder
}

func compare(a,b int) (string, error) {
	if (a > b) {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater!")
	}
}