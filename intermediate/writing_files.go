package main

import (
	// "bufio"
	"fmt"
	"os"
)


func main(){
	exampleFile, err := os.Create("intermediate/exampleFile.txt")
	if err != nil {
		fmt.Printf("Error in creating file:", err)
		return
	}
	fmt.Println("The file has been created successfully!")
	defer exampleFile.Close()

	var dataByte []byte = []byte("\tHello World! It's byte data\n\n\n\n\n")
	_ , err = exampleFile.Write(dataByte)
	if err != nil {
		fmt.Errorf("Error in creating file:", err)
		return
	}
	fmt.Println("The byte data has been written in to the file successfully!")

	// dataString := "\tHello World! This is string data!\n"	
	// _, err = exampleFile.WriteString(dataString)
	// if err != nil {
	// 	fmt.Errorf("Error in creating file:", err)
	// 	return
	// }
	// fmt.Println("The string data has been written in to the file successfully!")	

	// var dataByte1 []byte = []byte("\t\t\tHello World! Hello \n")
	// bufioWriter := bufio.NewWriter(exampleFile)
	// bufioWriter.Write(dataByte1)
	// bufioWriter.Flush()
	// fmt.Println("The byte data has been written in to the file successfully!")

	
}