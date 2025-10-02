package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {

	cmd := exec.Command("sleep", "5")
	err := cmd.Start()
	if err != nil {
		fmt.Println("error: ", err)
	}

	time.Sleep(2 * time.Second)

	err = cmd.Process.Kill()
	if err != nil {
		fmt.Println("error killing process : ", err)
	}

	fmt.Println("Process killed. ")
	err = cmd.Wait()
	if err != nil {
		fmt.Println("error waiting: ", err)
	}

	fmt.Println("Process finished.")

	// cmd := exec.Command("grep", "foo")
	// cmd.Stdin = strings.NewReader("foo is good\nbar\npistachio\ngiraffe")
	// o, _ := cmd.Output()
	// fmt.Println("The output is: ", string(o))

}
