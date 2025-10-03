package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func main(){
	fmt.Println("Process ID:", os.Getpid())
	// Creating a channel to receive signals 
	sigs := make(chan os.Signal, 1)

	// Using notify to send signals to the channel
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGUSR1)

	go func() {
		sig := <-sigs
		switch sig {
		case syscall.SIGINT :
			fmt.Println("Interupting signal.") //Ctrl + c
		case syscall.SIGTERM:
			fmt.Println("Terminating signal.") //kill -s SIGTERM {pid}
		case syscall.SIGHUP:
			fmt.Println("Hang up signal.") //kill -s SIGHUP {pid}
		case syscall.SIGUSR1:
			fmt.Println("SIGUSR1 signal.") //kill -s SIGUSER1 {pid}
			fmt.Println("User defined function is executed.")
		}
		fmt.Println("Graceful exit.")
		signal.Stop(sigs)
		os.Exit(0)
	
		// sig := <- sigs
		// fmt.Println("Interrupted: ", sig)
		// fmt.Println("Graceful shutdown.")
		// os.Exit(0)

	}()

	fmt.Println("Working...")

	for {
		time.Sleep(time.Second)
	}

	// signal.Stop(sigs)

	//kill -s SIGTERM {pid}

}