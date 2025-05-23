package main
 import "fmt"

func main(){
	/* Recover is a built-in function that is used to regain control of a panicing goroutine
	It's only useful inside DEFER functions 
	It's used to manage behavior of a panicing goroutine to avoid abrupt termination
	When recover uses in combination with deffer, recover can be used to handle or log errors and allow program to continue executing
	Recover function returns a value, But only when a panic occured.
	*/
	process ()
	fmt.Println("Returned from process!") 
	/* Output: 
	Start Process
	Recovered:  Something went wrong!
	Returned from process!
	*/

	/*
	Practical Use Cases:
	Recovery
	Cleanup
	Logging and Reporting

	Best Practices:
	Use with Defer
	Logging and reporting the panic (Avoid silent recovery)
	Avoid Overuse - Use ordinary error handling
	*/





}

func process () {
	defer func () {
		if r := recover();	r != nil { //If there is no panic, recover returns nil
			fmt.Println("Recovered:", r)
		} 
	}()
	fmt.Println("Start Process")
	panic("Something went wrong!")
	fmt.Println("End Process")
}