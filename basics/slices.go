package main

import (
	"fmt"
)


func main() {
	//A slice is an array but with uncertain size
	// var numbers []int 
	// var numbers1 []int = []int{1,2,3,4,5}
	// numbers2 := []int{9, 8, 7}

	slice := make([]int,5)

	a := [5]int{1, 2, 3, 4, 5}
	slice = a[1:4]

	fmt.Println(slice) //Prints [2,3,4]

	slice1 := append(slice, 6, 7)
	fmt.Println("Slice1:",slice1) //Prints [2 3 4 6 7]

	sliceCopy := make([]int,len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println("SliceCopy:", sliceCopy)

	//nil Slice
	// var nilSlice []int 

	for i, v := range(sliceCopy) {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	slice1[0] = 50
	// fmt.Println("First element of slice1:", slice1[0])

	// //Comparing two slices
	// fmt.Println("Are slice1 and sliceCopy the same?")

	// if slices.Equal(slice1, sliceCopy) {
	// 	fmt.Println("slice1 and sliceCopy are the same")
	
	// } else {
	// 	fmt.Println("slice1 and sliceCopy are NOT the same")
	// }

	// //Creating a multi-dimensional array using for loop
	// twoDSlice := make([][]int, 3)
	// fmt.Println("twoDSlice:", twoDSlice)
	// for i:=0; i<len(twoDSlice); i++ {
	// 	innerLength := i + 1
	// 	twoDSlice [i] = make([]int, innerLength)
	// 	for j:=0; j<innerLength; j++{
	// 		twoDSlice [i][j] = i + j
	// 	}

	// }
	// fmt.Println("twoDSlice:", twoDSlice)

	slice2 := make([]int, 7)
	slice2 = slice1[1:4]
	fmt.Println("The capacity of slice2 is:", cap(slice2))

	fmt.Printf("The slice1 is %v and capacity of it is: %v\n", slice1, cap(slice1))

	fmt.Printf("The slice2 is %v and capacity of it is: %v\n", slice2, cap(slice2))

	slice3 := slice2[1
	:5]
	fmt.Printf("The slice3 is %v and capacity of it is: %v\n", slice3, cap(slice3))












}