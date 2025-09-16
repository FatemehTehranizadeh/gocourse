package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func removeTempFilesandDirs(rootDir string, tempFileName string) {
	var dirsToDelete []string

	err := filepath.Walk(rootDir, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// fmt.Println("--------------- All Paths are:", path)
		if strings.HasPrefix(fileInfo.Name(), tempFileName) {
			// return os.RemoveAll(fileInfo.Name())
			if !fileInfo.IsDir() {
				fmt.Printf("Removing file: %s\n", path)
				return os.Remove(path)
			}

			if fileInfo.IsDir() {
				// fmt.Printf("Removing directory: %s\n", path)
				// return os.RemoveAll(path)
				dirsToDelete = append(dirsToDelete, path)
			}
		}
		return nil
	})
	checkError(err)

	for _, dir := range dirsToDelete {
		fmt.Printf("Removing directory: %s\n", dir)
		err := os.RemoveAll(dir)
		if err != nil {
			fmt.Printf("Error removing directory: %v\n", err)
		}
	}
}

func main() {

	rootTempDirectoryName := "rootTempDirectory"
	rootTempDir, err := os.MkdirTemp(".", rootTempDirectoryName)
	checkError(err)

	tempFileName := "tempFile"
	tempFile1, err := os.CreateTemp(filepath.Join(".", rootTempDir), tempFileName)
	checkError(err)
	fmt.Println("A new temporary file has been created:", tempFile1.Name())

	// Creating a temporary directory
	tempDirName := "tempDir"
	tempDir1, err := os.MkdirTemp(filepath.Join(".", rootTempDir), tempDirName)
	checkError(err)
	fmt.Println("A new temporary directory has been created:", tempDir1)

	tempFile2, err := os.CreateTemp(filepath.Join(".", tempDir1), tempFileName)
	checkError(err)
	fmt.Printf("A new temporary file has been created: %v and its directory is %v\n", tempFile2.Name(), filepath.Dir(tempFile2.Name()))

	defer removeTempFilesandDirs("/home/reera/gocourse", rootTempDirectoryName)
	defer removeTempFilesandDirs(filepath.Join(".", rootTempDir), tempFileName)
	defer removeTempFilesandDirs(filepath.Join(".", rootTempDir), tempDirName)
	defer tempFile1.Close()
	defer tempFile2.Close()
	defer func() {
		fmt.Println("The End!!!!!")
	}()
}
