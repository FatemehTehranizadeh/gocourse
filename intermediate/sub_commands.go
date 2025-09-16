package main

import (
	"flag"
	"fmt"
	"os"
)


func main(){

	subcommand1 := flag.NewFlagSet("FirstSubCommand", flag.ExitOnError)
	subcommand2 := flag.NewFlagSet("SecondSubCommand", flag.ExitOnError)

	fisrtFlag := subcommand1.String("name", "Reera", "This is used for name")
	secondFlag := subcommand2.Bool("isFemale", false, "This is used for gender")

	flag.Parse()
	fmt.Println(fisrtFlag)
	fmt.Println(secondFlag)

	for i, arg := range os.Args {
		fmt.Printf("Arg %d : %v\n", i, arg)
	}

	
}