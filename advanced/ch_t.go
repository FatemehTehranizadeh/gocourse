package main

import (
	"fmt"
	"time"
)

func getName(ch chan string) string {
	var name string
	fmt.Println("Please enter your name: ")
	fmt.Scanln(&name)
	ch <- name
	return name
}

func greeting(name string, ch chan string) {
	ch <- fmt.Sprintf("Welcome dear %s, Glad to have you here!\n", name)
}

func main() {

	nameChannel := make(chan string, 1)
	greetingChannel := make(chan string, 1)
	userName := getName(nameChannel)

	go func() {
		greeting(userName, greetingChannel)
	}()
	go func() {
		fmt.Println(<-greetingChannel)
	}()
	time.Sleep(100 * time.Millisecond)

}
