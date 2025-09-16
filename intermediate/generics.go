package main

import "fmt"

func swap[T any](a,b T) (T,T) {
	return b, a
}

//Let's create a stack
type Stack[T any] struct {
	stackSlice []T
}

//Let's create its methods
//First we should create a push method for our stack
func (s *Stack[T]) push (a T) {
	s.stackSlice = append (s.stackSlice, a)
	s.printStack()
}

//Pop method
func (s *Stack[T]) pop () (T, bool) {
	if len(s.stackSlice) == 0 {
		s.printStack()
		var zero T
		return zero, false
	} else {
		lastElement := s.stackSlice[len(s.stackSlice)-1]
		s.stackSlice = s.stackSlice[:len(s.stackSlice)-1]
		s.printStack()
		return lastElement, true
	}
}

//is our stack empty? 
func (s *Stack[T]) isEmpty () bool {
	return len(s.stackSlice) == 0
}

//Print method for our stack
func (s *Stack[T]) printStack () {
	if len(s.stackSlice) == 0 {		
		fmt.Println("Your stack is empty!!!!")
	} else {
		fmt.Println(s.stackSlice)
	}
}


func main(){

	// a, b := 2, 1
	// a, b = swap(a, b)
	// fmt.Printf("a is %v and b is %v\n", a, b) //a is 1 and b is 2

	s := Stack[float64]{}
	fmt.Println("Is the stack empty?", s.isEmpty())
	s.push(5656)
	s.push(87)
	s.push(12)
	s.push(0.45)
	s.push(-9)
	s.pop()
	fmt.Println("Is the stack empty?", s.isEmpty())

	stringStack := Stack[string]{}
	fmt.Println("Is the stack empty?", stringStack.isEmpty())
	stringStack.pop()
	stringStack.push("Hello")
	stringStack.push("World")
	stringStack.push("Jane!")
	fmt.Println("Is the stack empty?", stringStack.isEmpty())

	








}