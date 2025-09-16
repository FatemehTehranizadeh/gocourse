package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width  float64
	length float64
}

type RectangleNotGeometry struct {
	width  float64
	length float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.length
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r RectangleNotGeometry) area() float64 {
	return r.width * r.length
}

func (r Rectangle) perimeter() float64 {
	return (r.length + r.width) * 2
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

//There is no perimeter() function for RectangleNotGeometry struct

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("The area is:", g.area())
	fmt.Println("The perimeter is:", g.perimeter())
}

func main() {
	rect1 := Rectangle{
		width:  10,
		length: 20,
	}

	rect2 := RectangleNotGeometry{
		width:  1,
		length: 2,
	}

	// measure(rect2) => We got an error because the method perimeter() has not been declared for RectangleNoGeometry struct
	fmt.Println("The area of rect2 is:", rect2.area())
	measure(rect1)

	customizedPrinter(422,"hfddfsds0",0.556, -2, -0.00066, true)

}

func customizedPrinter (i ... interface{}) {
	for _, v := range i {
		fmt.Println("Let's print:",v)
	}
}