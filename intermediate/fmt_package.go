package main

import "fmt"


func main(){

	//Printing Functions
	// fmt.Print("Hello")
	// fmt.Print("World!")
	// fmt.Print(12,456)

	// fmt.Println("Hello")
	// fmt.Println("World!")
	// fmt.Println(12,456)

	// name := "John"
	// age := 25
	// fmt.Printf("Name: %s, Age: %d\n", name, age)
	// fmt.Printf("Binary: %b, Hexadecimal: %X\n", age, age)

	// /*Formatting Functions
	// It's done using spring functions
	// Actually spring functions concatenate the different types of data together. 
	// If we want to do it with "+" operator, we should convert data types to string
	// but using sprint there is no need to do that.
	// */
	// s := fmt.Sprint("Hello", "World!", 123, 456) //It concatenates the inputs
	// fmt.Print(s) //Prints HelloWorld!123 456

	// s =  fmt.Sprintln("Hello", "World!", 123, 456) //It adds space between every two words
	// fmt.Print(s) //Prints Hello World! 123 456
	// fmt.Print(s) //It will be printed in a new line

	// sf := fmt.Sprintf("Name: %s, Age: %d", name, age)
	// fmt.Println(sf)
	// fmt.Println(sf)

	/*Scanning Functions
	Scanln stops scanning at the new line but scan will repeatedly want the input and wouldn't be stopped
	Scanln will stop execution as soon as it encounters enter (a new line character)
	But scan will keep asking you for correct input, even if a new line is hitted, it doesn't stop.
	Scanf expects the inputs in the exact format that we have declared between the double quotes
	*/
	// var name string
	// var age int

	// fmt.Println("Enter you name and age:")
	// // fmt.Scan(&name, &age)
	// // fmt.Scanln(&name, &age) 
	// fmt.Scanf("%s %d", &name, &age)
	// fmt.Printf("Name: %s , Age: %d\n", name, age)

	/*Error Formatting Functions
	Errorf function
	*/
	err := checkAge(20)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive", age)
	}
	return nil
}