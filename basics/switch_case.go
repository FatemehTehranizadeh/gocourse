package main

import "fmt"


func main() {
	var score int = 72
	
	switch {
	case score >= 90: {
		fmt.Println("Grade A")
	}
	case score >= 80: {
		fmt.Println("Grade B")		
	}
	case score >= 70: {
		fmt.Println("Grade C")		
	}
	fallthrough
	case score < 70: {
		fmt.Println("Grade D")		
	}	
	}

	checkType(4536)
	checkType(2.1)
	checkType(true)

}

func checkType (x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer!")
	case float64:
		fmt.Println("It's a float!")
	case string:
		fmt.Println("It's a string!")
	default:
		fmt.Println("Uknown Type")		
	}
}