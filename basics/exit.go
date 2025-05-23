package main

import (
	"fmt"
	"os"
)


func main(){
	/*
	Os.exit is a function that terminates the program immediately
	Halting the execution of the program completely without defering any functions or performing any cleanup operations
	It takes an integer argument which represents the status code returned to the operating system
	*/

	defer fmt.Println("Defer statement")

	fmt.Println("Starting the main function...")
	os.Exit(1)
	fmt.Println("Ending the main function...")

	/*
	Practical Use Cases:
	Error Handling
	Termination Conditions
	Exit codes

	Best Practices:
	Avoid deferred actions
	Use meaningful status code (0 for success and non-zero for failures)
	Use it only when it's necessary
	*/



}