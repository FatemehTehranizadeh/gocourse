package main

import "fmt"


func main(){
	adderMain := adder()

	fmt.Println(adderMain())
	fmt.Println(adderMain())
	fmt.Println(adderMain())
	fmt.Println(adderMain())

	subtractor := func() func(int) int {
		countdown := 100
		return func (x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(subtractor(5))
	fmt.Println(subtractor(10))
	fmt.Println(subtractor(20))


	greeting := func(name string) func() {
		message := "Welcome %v to my GitHub!!! \n"
		// fmt.Println("Your name?")
		// fmt.Scanln(&name)
		return func() {
			fmt.Printf(message, name)
		}
	}

	var name string
	fmt.Println("Your name?")
	fmt.Scanln(&name)

	bobWelcome := greeting(name)
	bobWelcome()


}

func adder() func() int {
	i := 0
	fmt.Println("The previous value of i is:", i)
	return func () int {
		i++
		fmt.Println("i + 1:", i)
		return i
	}
}