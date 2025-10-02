package main

import (
	"fmt"
	"sync"
	"time"
)

type leakyBucket struct {
	capacity int           // Maximum capacity of the bucket (max tokens it can hold).
	leakRate time.Duration // How fast the tokens leak out of the bucket (time duration for each token to leak).
	tokens   int           // Current tokens in the bucket.
	lastLeak time.Time     // The last time the tokens were leaked.
	mu       sync.Mutex
}

func newLeakyBucket(capacity int, leakRate time.Duration) *leakyBucket {
	return &leakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		tokens:   capacity, //The bucket is full
		lastLeak: time.Now(),
	}
}

func (lb *leakyBucket) allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	now := time.Now()
	elapsedTime := now.Sub(lb.lastLeak)
	// tokensToAdd := int(elapsedTime / lb.leakRate) //We want to determine if we can accept new tokens or not, for example if we accept
	//new requests every 500 ms, if 200 ms has been passed, we can't add new tokens.
	// lb.tokens += tokensToAdd
	// lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate) //We should add time according to new tokens we add and the time they consumed

	if elapsedTime >= lb.leakRate {
		lb.tokens += 1
	}

	if lb.tokens > lb.capacity {
		lb.tokens = lb.capacity
	}

	lb.lastLeak = lb.lastLeak.Add(time.Duration(lb.leakRate))

	// fmt.Printf("Tokens added %d, Tokens subtracted %d, Total tokens: %d\n", tokensToAdd, 1, lb.tokens)
	// fmt.Printf("Last leak time: %v\n", lb.lastLeak)

	if lb.tokens > 0 {
		lb.tokens--
		return true
	}
	return false

}

func main() {
	fmt.Println("Start")
	lb := newLeakyBucket(3, time.Millisecond*500)

	var wg sync.WaitGroup
	var requests map[int]string
	requests = make(map[int]string)
	requestsLength := 20

	for i := 1; i <= requestsLength; i++ {
		requests[i] = fmt.Sprintf("The %dth request", i)
	}

	for i := 1; i <= requestsLength/2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if lb.allow() {
				fmt.Println("Request allowed: ", requests[i])
				delete(requests, i)
			} else {
				fmt.Println("Request denied: ", requests[i])
			}
			// time.Sleep(time.Millisecond * 800)
		}()
	}

	time.Sleep(time.Millisecond * 800)

	// for i := 11; i <= requestsLength; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		if lb.allow() {
	// 			fmt.Println("Request allowed: ", requests[i])
	// 			delete(requests, i)
	// 		} else {
	// 			fmt.Println("Request denied: ", requests[i])
	// 		}
	// 		time.Sleep(time.Millisecond * 200)
	// 	}()
	// }
	wg.Wait()

}
