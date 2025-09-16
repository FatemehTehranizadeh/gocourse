package main

import (
	"fmt"
	"os"
	"strings"
)
func checkError(err error){
	if err != nil{
		panic(err)
	}
}


func main(){
	user := os.Getenv("USER")
	// list := os.LookupEnv()

	fmt.Println("User value:", user)
	
	err := os.Setenv("FRUIT","BANANA")
	checkError(err)
	err = os.Unsetenv("FRUIT")
	checkError(err)	
	envVarsList := os.Environ()
	// fmt.Println("Environment variables list:", envVarsList)

	// fmt.Println(strings.SplitN(envVarsList[2], "=", 2))

	for _, envVar := range envVarsList {
		fmt.Println(strings.SplitN(envVar, "=", 2)[0])
	}

}