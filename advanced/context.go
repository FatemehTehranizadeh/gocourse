package main

import (
	"context"
	"fmt"
	"time"
)

/*
Contexts are instances of structs, It means that they are objects.
Context is a type from context package
Contexts are used to carry deadlines, cancellation signals, and request scooped values.
Contexts are closely associated with APIs. Contexts are frequently used when we are creating APIs.
Context is an object that carries information about deadlines, cancellation signals and
other request scooped values across API boundaries and goroutines. It is used to manage the life cycle of processes and
to signal when the operations should be aborted.
Difference between background context and Todo context: context.TODO is used when you are unsure about what
kind of context you wanat to use or you wanna change it later. It acts like a place holder.
*/

func checkEvenOdd(ctx context.Context, num int) string {
	select {
	case <-ctx.Done():
		return "Operation cancelled"

	default:
		if num%2 == 0 {
			return fmt.Sprintf("%d is even.\n", num)
		} else {
			return fmt.Sprintf("%d is odd.\n", num)
		}
	}
}

func main() {
	bkgctx := context.Background()

	result := checkEvenOdd(bkgctx, 5)
	fmt.Print(result)

	ctxWithTimeout, cancel := context.WithTimeout(bkgctx, time.Second * 2)
	defer cancel()

	result = checkEvenOdd(ctxWithTimeout, 7)
	fmt.Print(result)

	time.Sleep(3 * time.Second)
	result = checkEvenOdd(ctxWithTimeout, 7)
	fmt.Println(result)
}

// func main(){

// 	todoctx := context.TODO()
// 	bkgctx := context.Background()

// 	ctx1 := context.WithValue(todoctx, "name", "John")
// 	fmt.Println(ctx1)
// 	fmt.Println(ctx1.Value("name"))

// 	ctx2 := context.WithValue(bkgctx, "city", "New York")
// 	fmt.Println(ctx2)
// 	fmt.Println(ctx2.Value("city"))
// }
