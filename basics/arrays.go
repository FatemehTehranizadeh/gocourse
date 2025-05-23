package main

import "fmt"


func main() {
	// var arrayName [size]elementType

	// var numbers [5]int
	// fmt.Println(numbers)

	// numbers[0] = 500
	// numbers[len(numbers)-1] = 600
	// fmt.Println(numbers)

	// fruits := [4]string{"Apple","Banana","Orange","Grapes"}
	// fmt.Println(fruits)
	// fmt.Println(fruits[2])

	// originalArray := [3]int{1,2,3}
	// fmt.Println(originalArray)
	// copiedArray := originalArray 
	// copiedArray[0] = 100
	// originalArray[0] = 110


	// fmt.Println(originalArray)
	// fmt.Println(copiedArray)

	// for i:=0; i < len(numbers); i++ {
	// 	fmt.Printf("Element at index %d is %d \n",i,numbers[i])
	// }

	// for index, value := range numbers {
	// 	fmt.Printf("Element at index %d is %d \n",index,value)
	// }

	// //Blank Identifier => Underscore => Used to store unused variables
	// for _ , v := range numbers {
	// 	fmt.Printf("Value: %d\n", v)
	// }

	// //If we want to declare a value but not use it temporarily
	// b := 3
	// _ = b

	// //Comparing arrays
	// array1 := [3]int {1,2,3}
	// array2 := [3]int {1,2,3}
	// fmt.Println("Does array1 equal to array2?",array1==array2)

	// //Multidimentional array (Matrix)
	// var matrix [3][3]int = [3][3] int{
	// 	{1,2,3},
	// 	{4,5,6},
	// 	{7,8,9},
	// }
	// fmt.Println(matrix)


	//Pointers
	originalArray := [3]int{1,2,3}
	var copiedArray *[3]int

	copiedArray = &originalArray 

	copiedArray[0] = 100
	// originalArray[0] = 110


	// fmt.Println(originalArray)
	fmt.Println(*copiedArray)


	







}