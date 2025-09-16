package main

import (
	"fmt"
	"time"
)

func main() {
	greetingChannel := make(chan string)
	// greetingString := "hello"

	// sliceChannel := make(chan []int)
	// infiniteChannel := make(chan int)
	// errorChannel := make(chan error)
	// exampleSlice := make([]int,0)

	// go func ()  {
	// 	for i := 0; i <= 8; i++ {
	// 		exampleSlice = append(exampleSlice, i)
	// 	}
	// 	sliceChannel <- exampleSlice
	// }()

	// go func() {
	// 	i := 0
	// 	for true {
	// 		i++
	// 	}
	// 	infiniteChannel <- i
	// }()

	// go func() error  {
	// 	err := fmt.Errorf("There is an error:")
	// 	// errorChannel <- err
	// 	return err
	// }()

	// sliceReceiver := <- sliceChannel
	// fmt.Println(sliceReceiver)

	// infiniteReceiver := <- infiniteChannel
	// _ = infiniteReceiver

	// errorReceiver := <- errorChannel
	// fmt.Println(errorReceiver)

	go func() {
		for _ , e := range "abcdefg " {
			greetingChannel <- "Alphabet: " + string(e)
		}
	}()

	for true {
		receiver := <- greetingChannel
		fmt.Println(receiver)
		if receiver == "Alphabet: " + string(" "){
			break
		}
	}

	time.Sleep(50 * time.Millisecond)
	fmt.Println("End of program")

}
