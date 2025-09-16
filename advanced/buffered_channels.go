package main

import (
	"bytes"
	"fmt"
	"log"
	"time"
)

// func main() {
// 	// Buffered channels are blocked in two situations: Blocking on send if the capacity of buffer is full,
// 	// Blocking on recieve if the buffer is empty
// 	bufferedChannel := make(chan int, 2)
// 	bufferedChannel <- 1
// 	bufferedChannel <- 2
// 	fmt.Println(<-bufferedChannel)
// 	fmt.Println(<-bufferedChannel)
// 	bufferedChannel <- 3 //When channel is full, it needs a reciever to release. Then it can accept new values and we don't have error.
// 	fmt.Println(<-bufferedChannel) // Prints 3
// 	fmt.Println("Buffered Channels.")
// }

func main() {
	ch := make(chan int, 2)
	var logBuffer bytes.Buffer
	logger := log.New(&logBuffer, "", log.Ldate|log.Ltime)
	logger.SetPrefix("INFO: ")
	ch <- 1
	ch <- 2
	go func() {
		fmt.Println("Starting the goroutine")
		time.Sleep(2 * time.Second)
		fmt.Println("Ending the goroutine")
		fmt.Println("Received in goroutine: ", <-ch)
		logger.Println("Goroutine 1 is executing...")
	}()
	fmt.Println("Continuing the main goroutine")
	ch <- 3
	fmt.Println("Received: ", <-ch)
	fmt.Println("Received: ", <-ch)
	fmt.Println("Finish the main goroutine")
	// In this part of code we may have different outputs due to the scheduling of goroutines. For example we may see the
	// print statement "Received in goroutine:  1" at the end of the other print statements
	fmt.Println(logBuffer.String())
}
