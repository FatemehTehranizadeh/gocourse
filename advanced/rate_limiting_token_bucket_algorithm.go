package main

import (
	"fmt"
	"time"
)

type rateLimiter struct {
	tokens     chan struct{}
	refillTime time.Duration
}

func newRateLimiter(rateLimit int, refillTime time.Duration) *rateLimiter {
	rl := &rateLimiter{
		tokens:     make(chan struct{}, rateLimit),
		refillTime: refillTime,
	}

	for range rateLimit {
		rl.tokens <- struct{}{}
	}

	go rl.refillRateLimiter()
	return rl
}

func (rl *rateLimiter) refillRateLimiter() {
	ticker := time.NewTicker(rl.refillTime)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			select {
			case rl.tokens <- struct{}{}:
			default:
			}
		}
	}
}

func (rl *rateLimiter) allow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

func main() {

	rl := newRateLimiter(3, time.Second*1)

	for range 10 {
		if rl.allow() {
			fmt.Println("Request allowed.",time.Now())
		} else {
			fmt.Println("Request denied.", time.Now())
		}
		time.Sleep(time.Millisecond * 500)
	}

}
