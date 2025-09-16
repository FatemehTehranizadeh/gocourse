package main

import (
	"fmt"
	"time"
)


func main(){

	fmt.Println("The current time is:", time.Now())

	specificTime := time.Date(2024,time.December,2,14,32,46,0,time.UTC)
	fmt.Println("The specific time is:", specificTime)

	// parsedTime, _ := time.Parse("2006-01-02 15:04:05","2021-05-14 13:21:05") // Converting string to time.time
	parsedTime1, _ := time.Parse("06-1-2","21-5-4")
	fmt.Println("The parsed time is: ", parsedTime1)

	fmt.Printf("Formatted time is %v and its type is %T", time.Now().Format("2006-01-02 15:04:05"),time.Now().Format("2006-01-02"))

	//Manipulating time
	


}