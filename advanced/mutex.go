package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	var mux sync.Mutex

	numGoroutines := 5
	wg.Add(numGoroutines)

	increment := func() {
		defer wg.Done()
		for range 1000 {
			mux.Lock()
			counter++
			mux.Unlock()
		}
	}

	for range numGoroutines {
		go increment()
	}

	wg.Wait()
	fmt.Println("The final value is: ", counter)
}

// type counter struct {
// 	mux   sync.Mutex
// 	value int
// }

// func (c *counter) increment() {
// 	c.mux.Lock()
// 	defer c.mux.Unlock()
// 	c.value++
// }

// func (c *counter) read() int {
// 	c.mux.Lock()
// 	defer c.mux.Unlock()
// 	return c.value
// }

// func main() {
// 	var wg sync.WaitGroup

// 	c := &counter{}

// 	for i := 0; i < 15; i++ {
// 		wg.Add(1)
// 		go func() {
// 			for range 100 {
// 				c.increment()
// 			}
// 			defer wg.Done()
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println("The final value is: ", c.read())

// }
