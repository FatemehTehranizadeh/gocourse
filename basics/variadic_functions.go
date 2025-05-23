package main

import "fmt"


func main(){
	/* Variadic functions can accept an uncertain numbers of arguments
	The specification of it is ellipsis =. "..."
	func functionName (param1 type1, param2 type2, param3 ...type3) returnType {
	function body
	}
	Variadic parameter should be the last parameter is the function signature
	*/
	fmt.Println(sum(0,0,2,0,0,0,0,6))
	statement, total := sumWithString("The sum of 1, 2, 3 and 4: ",1,2,3,4)
	fmt.Println(statement, total)

	//Slices
	numbers := []int {1,2,3,4,5,9}
	sequence3, total3 := sum(3, numbers...)
	fmt.Println("Sequence 3:", sequence3, "Total3:", total3)





}
func sum(sequence int, nums ... int) (int,int) {
	total := 0
	for _ , v := range nums {
		total += v
	}
	return sequence, total
}

func sumWithString(returnString string, nums ... int) (string,int) {
	total := 0
	for _ , v := range nums {
		total += v
	}
	return returnString, total
}