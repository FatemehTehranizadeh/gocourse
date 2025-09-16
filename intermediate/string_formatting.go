package main

import "fmt"


func main() {
	// String formatting refers to the techniques used to create formatted outputs from variables or constants

	num := 74
	fmt.Printf("%08d\n", num)

	s := "Hello"
	fmt.Printf("|%16s|\n", s) // Numberas are right-aligned
	fmt.Printf("|%-16s|\n", s) // Numberas are left-aligned

	s1 := "Hello \nWorld!"
	s2 := `Hello \nWorld!`

	fmt.Println(s1)
	fmt.Println(s2)
}