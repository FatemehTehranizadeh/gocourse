package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		// time.Sleep(time.Second)
		ch1 <- 1
	}()

	go func() {
		ch2 <- 2
	}()

	for {
		select {
		case msg := <-ch1:

			fmt.Println("New message from Channel 1: ", msg)

		case msg := <-ch2:

			fmt.Println("New message from Channel 2: ", msg)

		case <-time.After(time.Second * 2):
			close(ch1)
			close(ch2)
			return

		default:
			fmt.Println("Default.")
			time.Sleep(time.Second)
		}
	}
}
