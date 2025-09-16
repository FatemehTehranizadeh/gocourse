package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0{
		return 0, errors.New("Math Error: Only positive numbers can be an argument of square root")
	} else {
		return 1, nil
	}
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error: The slice is empty!")
	}
	return nil
}


func main(){

	// result, err := sqrt(-1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("The result:", result)
	
	// if err := process([]byte{}); err != nil {
	// 	fmt.Println(err)
	// } 
	// fmt.Println("Data has been processed successfully!")

	// if err := eprocess(); err != nil {
	// 	fmt.Println(err)
	// } 

	if err := readData(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data read successfully.")

}

//Creating a custom error 

type myError struct {
	message string
}

//Using Error built-in function
func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func eprocess() error {
	return &myError{"Custom Error"}
}

func readConfig () error {
	return errors.New("Config Error!")
}

func readData () error {
	if err := readConfig(); err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	return nil
}