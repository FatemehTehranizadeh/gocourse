package main

import (
	"fmt"
	"time"
)


func main(){

	data := make(chan int, 2)
	quit := make(chan bool)

	go func() {
		for i := range 6 {
			data <- i
			time.Sleep(time.Second)
		}
		quit <- true
	}()


	for {
		select {
		case msg := <- data :
			fmt.Println("data is received: ", msg)

		case <- quit:
			fmt.Println("BYE.")
			close(data)
			close(quit)
			return

		default:
			fmt.Println("Waiting for a message...")
			time.Sleep(time.Millisecond * 500)
		}
	}
}