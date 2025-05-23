package main

import "fmt"


func main(){
	//Making a while loop using for loop
	// var i int = 1
	// for i<=5 {
	// 	fmt.Println("Iteration:", i)
	// 	i++
	// }

	sum := 0
	for {
		sum += 10
		fmt.Println("Sum:", sum)
		if sum >= 50 {
			break
		}
	}

}