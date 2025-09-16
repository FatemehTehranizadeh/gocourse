package main

import (
	"fmt";
	_ "embed"
)

//go:embed exampleFile.txt
var content string


func main() {

	fmt.Println("Embedded content:", content)
	fmt.Println("Embedded file from basics folder:", EmbeddedFile)



}