package main

import (
	"fmt"
	"time"
)

func main() {
	// Unbuffered channels need an immediate receiver and if there isn't, the app crashes.
	ch := make(chan int)
	
	go func() {
		fmt.Println(<-ch)
		time.Sleep(2 * time.Second)
	}()
	// ch <- 1
	fmt.Println("The rest of the main function")
}
