package main

import "fmt"


func main(){

	// //Simple iteration over a range
	// for i := 1; i<=5; i++ {
	// 	fmt.Println("i equals to:",i)

	// }

	// //Iteration over a collection
	// // numbers := []int {1,2,3,4,5,6}
	// var numbers []int = []int {1,2,3,4,5,6}
	// for index, value := range numbers {
	// 	fmt.Printf("Index is %v and Value is %d\n",index,value)
	// }
	// for i:= 1; i<=10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Printf("The i is %d\n",i)

	// 	if i == 8 {
	// 		break
	// 	}
		
	// }

	// rows := 5

	// for i :=1; i<=rows; i++ {
	// 	// inner loop for spaces before starts
	// 	for j:=1; j<=rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for k:= 1; k<=2*i-1; k++ {
	// 		fmt.Print(("*"))
	// 	}
	// 	fmt.Printf("\n") //Move to next line or use fmt.Println()
		
	// }
	// for i :=rows-1; i>0; i-- {
	// 	// inner loop for spaces before starts
	// 	for j:=1; j<=rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for k:= 1; k<=2*i-1; k++ {
	// 		fmt.Print(("*"))
	// 	}
	// 	fmt.Printf("\n") //Move to next line or use fmt.Println()
		
	// }

	// for i :=rows; i>0; i-- {
	// 	// inner loop for spaces before starts
	// 	for j:=rows; j>0; j-- {
	// 		fmt.Print(" ")
	// 	}
	// 	for k:= 1; k<=2*i-1; k++ {
	// 		fmt.Print(("*"))
	// 	}
	// 	fmt.Printf("\n") //Move to next line or use fmt.Println()
		
	// }

	// for i:=2; i<=rows; i++ {
	// 	// inner loop for spaces before starts
	// 	for j:=rows; j>0; j-- {
	// 		fmt.Print(" ")
	// 	}
	// 	for k:= 1; k<=2*i-1; k++ {
	// 		fmt.Print(("*"))
	// 	}
	// 	fmt.Printf("\n") //Move to next line or use fmt.Println()
		
	// }

	for i:= range 8 {
		fmt.Println(i)
	}

	 



}