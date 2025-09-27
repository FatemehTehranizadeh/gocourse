package main

import (
	"fmt"
	"sync"
	"time"
)

type rateLimiter struct {
	mux       sync.Mutex
	count     int
	limit     int
	window    time.Duration
	resetTime time.Time
}

func newRateLimiter(limit int, window time.Duration) *rateLimiter {
	return &rateLimiter{
		limit:  limit,
		window: window,
	}
}

func (rl *rateLimiter) allow() bool {
	rl.mux.Lock()
	defer rl.mux.Unlock()
	now := time.Now()

	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.window)
		rl.count = 0
	}
	if rl.count < rl.limit {
		rl.count++
		return true
	}
	return false
}

func main() {

	rl := newRateLimiter(2, 2*time.Second)

	var requests map[int]string
	requests = make(map[int]string)

	for i:=1; i<=10; i++ {
		requests[i] = fmt.Sprintf("The %dth request", i)
	}


	for k, _ := range requests {
		if rl.allow() {
			fmt.Println("Request allowed: ", requests[k])
			delete(requests, k)
		} else {
			fmt.Println("Request denied: ", requests[k])
		}
		time.Sleep(time.Millisecond * 500)
	}
}
