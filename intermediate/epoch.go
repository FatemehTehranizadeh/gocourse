package main

import (
	"fmt"
	"time"
)


func main(){

	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Current unix time is:", unixTime)
	
	t := time.Unix(unixTime, 0) //Converting Unix time to regular time format - 0 is nanosecond
	fmt.Println(t)

	fmt.Println("Time:", t.Format("2006-01-02"))


	
}