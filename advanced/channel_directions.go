package main

import (
	"fmt"
	"time"
)

// Producer - consumer pattern
// Defining a send-only or receive-only channel with make is not reasonable, just defining the channles
// unidirectional in the functions or goroutins make sense.

func main() {

	ch := make(chan int)

	go producer(ch)
	go consumer(ch)

	time.Sleep(time.Second)
}

func producer(ch chan<- int) {
	for i := 0; i <= 5; i++ {
		time.Sleep(time.Millisecond * 50)
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for range ch {
		fmt.Println("data is received.")
	}
}
