package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	//Generating a random number between 1 and 100
	target := random.Intn(100) + 1

	fmt.Println("Welcome to the game!")
	fmt.Println("Guess which number I've considered! Ha ha ha")

	var guess int
	for {
		fmt.Println("Give me your guess:")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("Congratulations! You guess is correct!")
			break
		} else if guess > target {
			fmt.Println("Choose a lower value!")
		} else {
			fmt.Println("Choose a higher value!")
		}
	}



}