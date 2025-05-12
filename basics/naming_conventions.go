package main
import "fmt"

type EmployeeGoogle struct{
	FirstName string
	LastName string
	Age int
}

type EmployeeApple struct{
	FirstName string
	LastName string
	Age int
}

func main(){
	/* PascalCase
	It's been typically used in Go for naming structs, interfaces or enums
	Eg. CalculateArea, UserInfo, NewHTTPRequest
	*/


	/*
	snake_case: lower case and separated by underscore
	Eg. user_id, first_name, httpe_request
	naming variables, constants and file names
	*/

	/*
	UPPERCASE
	It's been used just for constants!
	*/

	/*
	mixedCase
	Similar to PascalCase but it starts with a lower letter
	naming variables or certain identifiers
	Only use mixedCase or snake_case not a mix of them
	*/

	/*
	Package names should be in lower case and without underscore

	*/

	// var employeeID = 1001
	var employeeGoogle = EmployeeGoogle {FirstName: "Reera"}
	employeeGoogle.Age = 25
	fmt.Println("The name of Google employee is:",employeeGoogle.FirstName )

	p := EmployeeApple {FirstName: "Narges"}
	p.Age = 28
	fmt.Println("The Age of Apple emploee is:", p.Age)




	
}