package main

import "fmt"

var fiboNums = make(map[int]int)

func main() {

	fmt.Println(factoriel(5))
	fmt.Println(sumOfDigits(123))
	fmt.Println(fibonacci(5))
}

func factoriel(n int) int {
	if n == 0 {
		return 1
	}
	return n * factoriel(n-1)
}

func sumOfDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}

func fibonacci(n int) int {
	val, ok := fiboNums[n]
	if ok {
		return val
	}

	var result int
	if n == 0 {
		result = 0
	} else if n == 1 {
		result = 1
	} else {
		result = fibonacci(n-1) + fibonacci(n-2)
		// fmt.Printf("for n = %d the Fibo number is %d\n", n, currentFibNum)
	}
	fiboNums[n] = result
	fmt.Printf("for n = %d the Fibo number is %d\n", n, fiboNums[n])
	return result
}
