package main

import (
	"errors"
	"fmt"
)

type customError struct {
	code    int
	message string
	er      error
}

func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v", e.code, e.message, e.er)
}

// func doSomething () error {
// 	return &customError{
// 		code: 500,
// 		message: "Something went wrong",
// 	}
// }

func doSomething() error {
	if err := doSomethingElse(); err != nil {
		return &customError{
			code:    500,
			message: "Something went wrong",
			er: err,
		}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("internal error!")
}
func main() {

	if err := doSomething(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Something has been done successfully!")

}
