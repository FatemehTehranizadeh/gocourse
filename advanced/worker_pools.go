package main

import (
	"fmt"
	"sync"
)

/*
Managing a group of workers (goroutines) which are processing a queue of tasks in a channel

First thing in worker pools design pattrns is creating a task channel
Then we create worker goroutines and process the tasks
Next we dstribute the tasks; We send tasks to the task channel
The we implement gracefull shutdown
*/

func worker(wg *sync.WaitGroup, jobID int, tasksCh <-chan int, resultsCh chan<- int) {
	for t := range tasksCh {
		wg.Add(1)
		jobID ++
		fmt.Printf("Processing the task %d with id %d\n", t, jobID)
		go func(t int) {
			resultsCh <- t * 2
			defer wg.Done()
		}(t)
	}
}

const (
	numWorker = 3
	numJobs   = 10
)

func main() {
	jobsCh := make(chan int, numJobs)
	resultsCh := make(chan int, numJobs)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(jobsCh)
		for i := range numJobs {
			jobsCh <- i
		}
	}()

	worker(&wg, 0, jobsCh, resultsCh)

	wg.Add(1)
	go func ()  {
		close(resultsCh)
		wg.Done()
	}()

	wg.Add(1)
	go func ()  {
		for i := range resultsCh {
			fmt.Println("Results: ", i)
		}
		defer wg.Done()
	}()

	wg.Wait()

}
