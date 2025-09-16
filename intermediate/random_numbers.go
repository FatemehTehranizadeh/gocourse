package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	// "time"
)


func main(){

	val := rand.New(rand.NewSource(23))
	// val := rand.New(rand.NewSource(time.Now().Unix()))

	fmt.Println(val.Intn(100) + 1)
	fmt.Printf("%.5f",val.Float64())

	for {
		// Show the menu
		fmt.Println("Welcome to the Dice Game!")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice (1 or 2): ")

		var choice string
		reader := bufio.NewReader(os.Stdin)
		choice,_ = reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "2" {
			fmt.Println("\n---------  Have a good day! -----------\n")
			return
		} else if choice !="1" && choice !="2" {
			fmt.Println("\n !!!!!!!!!!! You've entered a wrong number, Try Again!\n !!!!!!!!!!!!!")
			continue
		}

		diceRandSeed := rand.New(rand.NewSource(time.Now().Unix()))
		diceNumber := diceRandSeed.Intn(12) + 1
		fmt.Println("\n ##### The dice is:", diceNumber, "######\n")

		fmt.Println("\n ******** Do you want to continue? (y/n) ********** \n")
		
		var rollAgain string
		fmt.Scanln(&rollAgain)

		if rollAgain == "n" {
			fmt.Println("\n ---------- Thanks for playing!------------ \n")
			return
		} else if rollAgain != "n" && rollAgain != "y" {
			fmt.Println("\n!!!!!!!!! Invalid input !!!!!!!!\n")
			continue
		}
	}
}