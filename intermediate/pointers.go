package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/* Pointers
	A variable that stores the memory address of another variable
	We pass large data structures efficiently between functions using pointers
	Pointers have their specific types just like data structures
	We can modify the value of another variable indirectly using pointers
	Declaration: var ptr *string
	Referencing =? &
	Dereferencing => *
	We can access struct fields through pointers
	*/

	var ptr *string
	var name1 string = "Reera"
	var name2 string

	ptr = &name1
	name1 = "Mohammad"
	name2 = *ptr

	fmt.Println("name1: ", name1)
	fmt.Println("name2: ", name2)
	fmt.Println("Pointer Value: ", *ptr)

	// When we use a function which it takes a pointer as the input, it changes the original value of the variable, not a copy of the variable
	var a int = 10
	var b int = 20

	modifyValueByPointer(&a)
	modifyValueByPointer(&a)
	modifyValueByPointer(&a)
	modifyValueByPointer(&a)
	fmt.Println(modifyValueWithoutPointr(b))
	fmt.Println("The value of a:", a) //14
	fmt.Println("The value of b:", b) //20

	//Using unsafe package for pointer
	var x *int
	res := unsafe.Pointer(&x)
	fmt.Println(res)

}

func modifyValueByPointer(ptr *int) {
	*ptr++
}

func modifyValueWithoutPointr(a int) int {
	b := a + 1
	return b
}
