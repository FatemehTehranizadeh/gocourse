package main

import "fmt"

type Rectangle struct {
	width float64
	length float64
}

func (r Rectangle) area () float64 {
	return r.length * r.width
}

func (r *Rectangle) scale (factor float64) {
	r.length *= factor
	r.width *= factor
}

type customizeInt int 
func (i customizeInt) isPositive () bool {
	return i > 0
}

type Shape struct {
	Rectangle //Anonymous field
	r Rectangle //Known field
}


func main() {
	rect := Rectangle {
		width: 10,
		length: 20,
	}
	fmt.Println(rect.area())
	rect.scale(0.5)
	fmt.Println(rect.area())

	//Methods can be associated with any type
	myInt := customizeInt(2)
	fmt.Printf("Is %d Positive? %v\n",myInt, myInt.isPositive())

	s := Shape {
		Rectangle: Rectangle{
			width: 10,
			length: 20,
		},
	}
	fmt.Println("The area is:", s.area()) //We can access the methods of Rectangle struct from Shape when we define an anonymous field


}