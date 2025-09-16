package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// func main() {

// 	counter := true

// 	go func() {
// 		if counter {
// 			fmt.Println("I'm the goroutine 1!", time.Now())
// 		}
// 		counter = false
// 	}()

// 	go func() {
// 		fmt.Println("Starting the goroutine 2...", time.Now())
// 		time.Sleep(time.Millisecond * 100)
// 		if !counter {
// 			fmt.Println("I'm the goroutine 2!", time.Now())
// 		}
// 		counter = true
// 	}()

// 	time.Sleep(time.Second)
// 	fmt.Println("========== The End ==========", time.Now())

// }

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)
	var res float64

	go func() {
		fmt.Println("I'm the goroutine 1!", time.Now())
		res = 2.2 + 3.3
		fmt.Println("Value of result in goroutine 1 is:", res)
		// fmt.Println(time.Second)
		wg.Done()
	}()

	go func() {
		fmt.Println("Starting the goroutine 2...", time.Now())
		fmt.Println("I'm the goroutine 2!", time.Now())
		res = 1.1 + 2.2
		fmt.Println("Value of result in goroutine 2 is:", res)
		fmt.Println(os.Args)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("========== The End ==========", time.Now())
}
