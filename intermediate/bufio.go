package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main(){
	
	bufIOReader := bufio.NewReader(strings.NewReader("Hello, Bufio Packageee!\n How are you doing?"))
	dataPlace := make([]byte, 20)
	nBytesRead, err := bufIOReader.Read(dataPlace)
	if err != nil {
		fmt.Println("Error in Reading:", err)
		return
	}
	fmt.Printf("%d bytes are read and the content is %s\n", nBytesRead,dataPlace[:nBytesRead])

	stringRead, err := bufIOReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error in Reading:", err)
		return
	}
	fmt.Printf("The content is %s\n", stringRead)

	bufIOWriter := bufio.NewWriter(os.Stdout)
	var dataToWrite []byte = []byte ("Hello, Bufio Package!\nHow are you doing?\n")
	nBytesWrote, err := bufIOWriter.Write(dataToWrite)
	if err != nil {
		fmt.Println("Error in Writing:", err)
		return
	}
	fmt.Printf("%d bytes are written.\n", nBytesWrote)

	// Flush the buffer to ensure all data is written to os.Stdout
	err = bufIOWriter.Flush() // It prints what it wrote
	if err != nil {
		fmt.Println("Error in Writing:", err)
		return
	}

	nBytesWrote, err = bufIOWriter.WriteString("This is an example string\n")
	if err != nil {
		fmt.Println("Error in Writing:", err)
		return
	}
	fmt.Printf("%d bytes are written.\n", nBytesWrote)
	bufIOWriter.Flush()

}