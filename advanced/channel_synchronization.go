package main

import (
	"fmt"
	"time"
)

// func main(){

// 	done := make(chan struct{})

// 	go func() {
// 		fmt.Println("Working...")
// 		time.Sleep(time.Second * 2)
// 		done <- struct{}{}
// 	}()

// 	<- done
// 	fmt.Println("Finished.")

// }

// func main() {
// 	numGoroutines := 3
// 	done := make(chan int)

// 	for i := 0; i <= numGoroutines; i++ {
// 		time.Sleep(time.Second * 2)
// 		go func(j int) {
// 			fmt.Printf("Goroutine %d is working...\n", j)
// 			// time.Sleep(time.Second)
// 			done <- j
// 		}(i)
// 	}

// 	for range numGoroutines - 1 {
// 		<-done
// 	}

// 	fmt.Println("All goroutines are done!")
// 	defer close(done)

// }

func main() {
	dataCh := make(chan string)

	go func() {
		for i := 0; i <= 5; i++ {
			dataCh <- fmt.Sprintln("data: ", i)
		}
		close(dataCh)
	}()

	for v := range dataCh {
		fmt.Printf("%v received. %v\n", v, time.Now().Format("15:04:05.000000"))
	}

}
