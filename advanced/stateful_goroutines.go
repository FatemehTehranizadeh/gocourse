package main

import (
	"fmt"
	"time"
)

type statefulWorker struct {
	value int
	ch    chan int
}

func (sw *statefulWorker) start() {
	go func() {
		for {
			select {
			case v := <-sw.ch:
				sw.value += v
				fmt.Println("The value: ", sw.value)
			}
		}

	}()
}

func (sw *statefulWorker) send(v int) {
	sw.ch <- v
}

func main() {
	sw := statefulWorker{ch: make(chan int)}
	sw.start()

	for i := 0; i <= 5; i++ {
		sw.send(i)
		time.Sleep(time.Millisecond * 700)
	}

}
