package main

import "fmt"


func main(){
	/* Panic is a built-in function that stops normal execution of a function immediately
	After stopping the normal process, the defer functions run
	It used to signal an unexpected error condition
	Its syntax:
	panic (interface {})
	Interface means that any value with any type can be used as the argument
	*/
	//Example of a valid input
	// process(10)

	//Example of an invalid input
	process(-3)




}

func process (input int) {
	defer fmt.Println("Deferred statement 1")
	defer fmt.Println("Deferred statement 2")
	if input < 0{
		defer fmt.Println("before Panic")
		panic("Input must be a non-negative value")
		// defer fmt.Println("after Panic")
	} 
	fmt.Println("processing input:", input)
}