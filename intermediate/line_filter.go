package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	exampleFile, err := os.Open("intermediate/exampleFile.txt")
	if err != nil {
		fmt.Printf("Error in opening the file:", err)
		return
	}
	fmt.Println("The file has been opened successfully!")
	defer func() {
		fmt.Println("The file is closing...")
		exampleFile.Close()
		fmt.Println("The file closed.")
	}()

	// keyword := make([]byte,20)
	// fmt.Println("Please enter the keyword:")
	// _, err = io.ReadFull(os.Stdin, keyword)
	// if err != nil {
	// 	fmt.Printf("Error in reading the keyword:", err)
	// 	return
	// }
	// fmt.Println("The keyword has been entered successfully and its value is:", keyword)

	keywordReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the keyword:")
	keyword, err := keywordReader.ReadString('\n')
	keyword = strings.TrimSpace(keyword)
	if err != nil {
		fmt.Printf("Error in reading the keyword:", err)
		return
	}
	fmt.Println("The keyword has been entered successfully and its value is:", keyword, "\n#############")

	scanner := bufio.NewScanner(exampleFile)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println("line:", line)
		if strings.Contains(line, keyword) {
			fmt.Println("Filtered line:",line)
			updatedLine := strings.ReplaceAll(line, keyword, "Kotlet")
			fmt.Println("Updated line:",updatedLine)
		}
	}





	


}
