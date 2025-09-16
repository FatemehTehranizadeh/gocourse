package main

import "fmt"


func main() {
	//Range keyword uses for iterating over data structures like arrays
	//Iterating over channels is common
	//Iterating over strings

	message := "Hello World"

	for i, v := range message {
		// fmt.Println("index: ", i, "Value: ", v) //Prints the unicode values
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}
	//Range works on a copy of the variable
	//Range iterates over arrays and slices and strings in order from the first element to the last one
	//For maps, it iterates over key-value pairs in a unspcified order
	//For channels, it iterates until the channel is closed.


}