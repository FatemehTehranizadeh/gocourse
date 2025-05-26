package main

import "fmt"

func init(){
	fmt.Println(("Initializing package 1..."))
}

func init(){
	fmt.Println(("Initializing package 2..."))
}

func init(){
	fmt.Println(("Initializing package 3..."))
}

func init(){
	fmt.Println(("Initializing package 4..."))
}


func main(){
	/* init is used for initialization tasks for the package before it is used
	 init function has no input and output
	 Go runs it automatically when the package is used
	 It happens before the main fucntion is executed
	 It just runs once
	 If there are multiple init function, they execute sequentially
	 init function is commonly used for initializing variables, performing setup operations,
	 registering component or configurations and initializing state required for the package 
	 to function correctly.

	*/
	fmt.Println("Inside the main function")

	/*
	Practical Use Cases:
	Setting up the tasks such as initializing global variables or constants required for the package
	Reading configuration files or environment variables
	Registering components or plugins with other parts of package or external systems
	Open datatbase connections or perform schema migrations needed by package
	
	Best Practices:
	Be careful about side effects of init function, It might effect the behavior of other packages or cause unexpected behavior
	Be careful about initialization order
	Document the purpose and side effects of init function
	*/
}