package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	joinedPath := filepath.Join("/","home", "reera", "downloads", "fileName.txt")
	fmt.Println("Result of join:", joinedPath)

	fmt.Println("Is joinedPath absolute?", filepath.IsAbs(joinedPath))
	fmt.Println("File name:", filepath.Base(joinedPath))
	fmt.Println("Directory of the file:", filepath.Dir(joinedPath))
	dir, file := filepath.Split(joinedPath)
	fmt.Printf("Directory of the file: %v \nThe name of the file: %v\n", dir, file)

	fmt.Println("The extension of file is:", filepath.Ext(file))
	fmt.Println("The name of file without its extension:",strings.TrimSuffix(file, filepath.Ext(file)))

	relativeDir, _ := filepath.Rel("a/b/c", "e/d/download/file.txt")
	fmt.Println("Relative directory:",relativeDir)
}
