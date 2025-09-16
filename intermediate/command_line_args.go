package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Command:", os.Args[0])

	for i, arg := range os.Args {
		fmt.Printf("Arg %d : %v\n", i, arg)
	}

	var nameVar string
	var ageVar int
	var isFemaleVar bool

	flag.StringVar(&nameVar, "name", "Reera", "This flag specifies the name")
	flag.IntVar(&ageVar, "age", 25, "This flag specifies the age")
	flag.BoolVar(&isFemaleVar, "isFemale", false, "This flag specifies the gender")

	flag.Parse()

	fmt.Println("Name:", nameVar)
	fmt.Println("Age:", ageVar)
	fmt.Println("isFemale:", isFemaleVar)

}
