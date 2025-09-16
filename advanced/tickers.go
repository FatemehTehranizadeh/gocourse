package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	// i := 0
	// for range 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// Running a task periodically
	// for {
	// 	select {
	// 	case <- ticker.C:
	// 		periodicTask()
	// 	}
	// }

	//Stopping a ticker gracefully
	// stop := time.After(4 * time.Second)

	// for {
	// 	select {
	// 	case <-stop:
	// 		fmt.Println("Tick is stopped.")
	// 		return
	// 	case <-ticker.C:
	// 		fmt.Println("Tick")
	// 		// default:
	// 		// 	fmt.Println("Waiting...")
	// 	}

	// }


}

func periodicTask() {
	fmt.Println("Performing a periodic task")
}
