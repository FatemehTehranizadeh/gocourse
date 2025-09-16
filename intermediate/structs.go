package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	married   bool
	isFemale  bool
	id        int
	address   Address
	PhoneNumbers
}

type Address struct {
	city    string
	country string
}

type Car struct {
	model  string
	year   int
	color  string
	number []int
}
type PhoneNumbers struct {
	homePhone string
	workPhone string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p Person) sex() string {
	if p.isFemale {
		return "Female"
	} else {
		return "Male"
	}
}

func (p *Person) incrementAgeByOne() {
	p.age++
		
}

func main() {
	/* Structs
	Composite data types that allow you to group different data types of variables under a single name
	Similar to classes, but they are lightweight and don't support inheritance
	We have "Anonymous structs" => Useful for temporary data structures

	*/

	p1 := Person{
		firstName: "John",
		lastName:  "Doe",
		address: Address{
			city:    "London",
			country: "England",
		},
		PhoneNumbers: PhoneNumbers{
			homePhone: "4564542",
			workPhone: "99888879",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.firstName)
	p1.id = 1

	c1 := Car{
		model: "Peugeot",
		color: "white",
		year:  2018,
	}

	fmt.Println(c1.year)

	//Anonymous Struct

	user := struct {
		userName string
		age      int
	}{
		userName: "user1234",
		age:      20,
	}
	fmt.Println(user.age)

	fmt.Printf("The full name of the Person %d is %s\n", p1.id, p1.fullName())
	fmt.Printf("The sex of the Person %d is %s\n", p1.id, p1.sex())

	//Creating instances of Person
	var people []Person
	var firstname, city string
	var age, idCounter int
	for {
		fmt.Println("Enter your name:")
		fmt.Scanln(&firstname)
		if firstname == "exit" {
			break
		}
		fmt.Println("Enter your age:")
		fmt.Scanln(&age)

		fmt.Println("Enter your city:")
		fmt.Scanln(&city)

		idCounter++

		p := Person{
			firstName: firstname,
			age:       age,
			address: Address{
				city: city,
			},
			id: idCounter,
		}
		people = append(people, p)

		fmt.Printf("A new person has been added with the name %s, city %s, age %d\n", p.firstName, p.address.city, p.age)
	}

	fmt.Println("All added persons are:")
	for _, v := range people {
		fmt.Printf("Name: %s,City: %s, Age: %d\n", v.firstName, v.address.city, v.age)
	}

	//Structs with anonymous fields
	fmt.Println(p1.homePhone) //We can access the properties of anonymous fileds directly


}
