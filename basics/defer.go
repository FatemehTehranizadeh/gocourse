package main

import "fmt"


func main(){
	/*Defer is a mechanism that allows you to postpone the execution of a function until the surrounding function returns
	It's useful for cleanup at the end of running code
	Its mechanism is Last-In First-Out
	It evaluates as soon as it's encountered
	*/

	process(1)

	/* Practical Uses of defer:
	Making APIs
	Making microservices
	Ensuring that resources like files or database connections are closed after they are opened
	Ensuring that a mutex is unlocked even if a function panics
	Logging and tracing entry and exit points of a function
	*/

	/*
	Best practices:
	Not use it in loops or nested functions because it can lead to subtle bugs if not handle carefully
	Use short and simple commands with differ
	*/
}
func process (i int) {
	defer fmt.Println("Deferred i value:", i)
	defer fmt.Println("First Deferred statement executed")
	defer fmt.Println("Second Deferred statement executed")
	defer fmt.Println("Third Deferred statement executed")
	defer fmt.Println("Fourth Deferred statement executed")
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("The actual value of i:",i)
}