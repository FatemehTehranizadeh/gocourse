package main

import (
	"fmt"
	"sync"
	"time"
)

// ٍExample with constuction
type worker struct {
	id   int
	wg   *sync.WaitGroup
	task string
}

func (w *worker) performTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("WorkerID %d started %s\n", w.id, w.task)
	time.Sleep(time.Second)
	fmt.Printf("WorkerID %d finished %s\n", w.id, w.task)
}

func main() {
	var wg sync.WaitGroup
	tasks := []string{"digging", "laying bricks", "painting"}

	for i, t := range tasks {
		wg.Add(1)
		w := worker{id: i, wg: &wg, task: t}
		go w.performTask(&wg)
	}

	wg.Wait()
	fmt.Println("All tasks are done.")
}

// Example using channels
// func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker %d is running...\n", id)
// 	time.Sleep(time.Second)
// 	for i := range tasks {
// 		results <- i * 2
// 	}
// 	fmt.Printf("Worker %d finished.\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 1
// 	numJobs := 5
// 	results := make(chan int, numJobs)
// 	tasks := make(chan int, numJobs)

// 	wg.Add(numWorkers)
// 	for i := range numWorkers {
// 		go worker(i, tasks, results, &wg)
// 	}

// 	for i := range numJobs {
// 		tasks <- i + 1
// 	}
// 	close(tasks)

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	for res := range results {
// 		fmt.Println("Result:", res)
// 	}

// }

// Example without using channels
// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker %d is working...\n", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("Worker %d finished.\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	wg.Add(numWorkers)

// 	go func() {
// 		for i := range numWorkers {
// 			worker(i, &wg)
// 		}
// 	}()

// 	// for i := range numWorkers{
// 	// 	go func(id int){
// 	// 		worker(id, &wg)
// 	// 	}(i)
// 	// }

// 	wg.Wait()
// }
