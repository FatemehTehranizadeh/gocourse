package main

import "fmt"


func main(){
	/* func <name> (input parameters list)(output parameters) or returnType {
	code 
	return value (If we mentioned a returnType)
	} */

	/* Rules: 
	Public functions should start with uppercase letters
	Private functions should start with lowercase letters
	
	zero or more parameters can be defiend with name and type
	parameters act as variable inside the function

	returnType specifies the type of values returned by the function
	The function can return multiple values too!

	In the function body, there is code and the last thing is return
	*/

	// fmt.Println(add(1,-3))

	// /* Anonymous functions or closures or function literals
	// Functions without a name explicity
	// */

	// func(){
	// 	fmt.Println("Hello World")
	// }()

	// greet := func(){
	// 	fmt.Println("Hello World")
	// }

	// greet()

	// //Taking functions as types
	// operation := add 
	// result := operation (3,5)
	// fmt.Println(result)

	//Passing a function as an argument
	result := applyOperation(5, 3, add)
	fmt.Println("result is:", result)

	//Return and using a function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("7 * 2:", multiplyBy2(7))

	//Functions are a first class citizens
	
}

func add(a,b int) int {
	return a + b
}

// A function that takes another function as an argument
func applyOperation (x int, y int, operation func(int, int) int) int {
	return operation(x,y)
}

func createMultiplier (factor int) func(int) int {
	return func (x int) int {
		return x * factor
	}
}
