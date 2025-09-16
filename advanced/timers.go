package main

import (
	"fmt"
	"time"
)

//TImer is not blocking but time.Sleep is.

func main() {
	timer := time.NewTimer(2 * time.Second)
	timer.Stop()
	defer timer.Stop()
	if timer.Stop() {
		fmt.Println("Timer is stopped.")
	}
	fmt.Println("Before reset")
	timer.Reset(1 * time.Second)
	fmt.Println("After reset")
	<-timer.C
	fmt.Println("Timeout.")

	done := make(chan struct{})

	go func() {
		longRunningOp()
		done <- struct{}{}
	}()

	select {
	case <-time.After(time.Second * 3):
		fmt.Println("Time out.")

	case <- done :
		fmt.Println("Task is completed.")
	}

}

func longRunningOp() {
	for i := range 20 {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
