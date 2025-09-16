
package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

type Employee struct {
	Person
	salary int
	empID  int
}

func (p Person) introduce() {
	fmt.Printf("Hi everyone! I'm %s %s. I'm %d years olddddd!\n", p.firstName, p.lastName,p.age)
}

//Overwriting a function
func (e Employee) introduce() {
	fmt.Printf("Hi everyone! I'm %s %s. I'm %d years old!. I earn %d every month\n", e.firstName, e.lastName, e.age, e.salary)
}

func main() {

	emp := Employee {
		Person: Person{
			firstName: "Reera",
			lastName: "Tehranizadeh",
			age: 25,
		},
		salary: 5000,
		empID: 1,
	}

	fmt.Println(emp)
	emp.introduce()

}

// func sumOfData (i...interface{}) float64 {
// 	total := 0.0
// 	for _, v := range i {
// 		total += v
// 	}
// 	return v
// }
