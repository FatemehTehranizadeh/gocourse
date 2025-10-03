package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string
	age  int
}

// ==================Working with methods
type Greeter struct{}

func (g Greeter) Greet(firstName, lastName string) string {
	return "Hello " + firstName + " " + lastName
}

func structFields(itf interface{}) {
	t := reflect.TypeOf(itf)
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumMethod(); i++ {
			method := t.Method(i)
			fmt.Printf("Method %d: %s\n", i, method.Name)
		}
	} else {
		fmt.Printf("The input is not struct. It's %T\n", itf)
	}

}

func main() {
	g := Greeter{}
	// t := reflect.TypeOf(g)
	v := reflect.ValueOf(g)
	// var method reflect.Method

	structFields(g)

	m := v.MethodByName("Greet")
	res := m.Call([]reflect.Value{reflect.ValueOf("Alice"), reflect.ValueOf("Doe")})
	fmt.Println(res[0])

}

//==================Working with Structs and Fields

// func main(){

// 		p := person{Name: "Alice", age:13}
// 	v := reflect.ValueOf(p)
// 	v1 := reflect.ValueOf(&p).Elem()

// 	for i:=0; i<v.NumField(); i++ {
// 		fmt.Printf("Filed %d: %v\n", i, v.Field(i))
// 	}

// 	nameField := v1.FieldByName("Name")
// 	fmt.Println("name field is: ", nameField)
// 	if nameField.CanSet() {
// 		nameField.SetString("Jane")
// 	} else {
// 		fmt.Println("It can't be set.")
// 	}

// 	fmt.Println("Modified Person: ", p)

// 	// var a int
// 	// a = 42
// 	// b := reflect.ValueOf(a)
// 	// fmt.Println("Value: ", b)
// 	// fmt.Println("Type: ", b.Type())
// 	// fmt.Println("Kind: ", b.Kind())
// 	// fmt.Println("Is int? ", b.Kind()==reflect.Int)
// 	// fmt.Println("Is string? ", b.Kind()==reflect.String)
// 	// fmt.Println("Is zero? ", b.IsZero())

// 	// c := 10
// 	// d := reflect.ValueOf(&c).Elem()
// 	// e := reflect.ValueOf(&c)
// 	// fmt.Println("Type of e: ", e.Type())
// 	// fmt.Println("Type of d: ", d.Type())
// 	// fmt.Println("Original value of c: ", d.Int())

// 	// d.SetInt(18)
// 	// fmt.Println("Modified value of c: ", d.Int())

// 	// var itf interface{}
// 	// itf = 1

// 	// v1 := reflect.ValueOf(itf)
// 	// fmt.Println("Value of itf: ", v1)
// 	// fmt.Println("Type of itf: ", v1.Type())
// 	// switch v1.Kind() {
// 	// 	case reflect.String :
// 	// 		fmt.Println("String value: ", v1.String())
// 	// 	case reflect.Int :
// 	// 		fmt.Println("Integer value: ", v1.Int())
// 	// }
// }
