package main

import "fmt"


func main() {

	var temperature int = 25

	if temperature >= 30 {
		fmt.Println("It's hot outside!")
	} else {
		fmt.Println("It's cool outside!")
	}

	var score int = 73
	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}

}