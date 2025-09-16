package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main(){
	// checkError(os.Mkdir("./subdir1", 0755))
	// checkError(os.WriteFile("./subdir1/file.txt", []byte(""), 0755))
	// checkError(os.MkdirAll("./subdir1/parent/child", 0755))
	// checkError(os.Mkdir("./subdir1/parent/child1", 0755))
	// checkError(os.Mkdir("./subdir1/parent/child2", 0755))
	// checkError(os.WriteFile("./subdir1/parent/child2/file.txt", []byte(""), 0755))

	// directories, err := os.ReadDir("subdir1/parent")
	// checkError(err)

	// for _ , dir := range directories {
	// 	fmt.Println(dir)
	// }

	checkError(os.Chdir("subdir1/parent"))
	wd, err := os.Getwd()
	checkError(err)
	fmt.Println("Working directory:", wd)

	fmt.Println("Let's walk through directories using filepath.WalkDir method!\nLet's go ~~~~~~~~~~~~")
	filepath.WalkDir("child2", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error during walking the directories:", err)
		}
		fmt.Println(path)
		return nil		
	})

	defer os.RemoveAll("/home/reera/gocourse/subdir1")



}