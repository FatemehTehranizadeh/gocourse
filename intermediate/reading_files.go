package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	exampleFile, err := os.Open("intermediate/exampleFile.txt")
	if err != nil {
		fmt.Printf("Error in opening the file:", err)
		return
	}
	fmt.Println("The file has been opened successfully!")
	defer exampleFile.Close()

	// reading the file 
	// readData := make([]byte,1024)
	// _, err = exampleFile.Read(readData)
	// if err != nil {
	// 	fmt.Printf("Error in reading the file:", err)
	// 	return
	// }
	// fmt.Println("The data has been read and It's ", string(readData))

	// Reading using Scanner

	scanner := bufio.NewScanner(exampleFile)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Printf("Error in scanning the file:", err)
		return
	}

}
