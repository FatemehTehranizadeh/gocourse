package main

import "fmt"


func main(){

	var a int = 3 
	b := float32(a)
	c := int16(a)
	
	var d float64 = 3.14
	e := int32(d)

	f := "Hello"
	g := []byte(f)

	h := []byte{300, 100}


	fmt.Println("a:", a,"\nb = float32(a): ", b, "\nc = int16(a): ", c, "\nd:", d, "\ne = int32(d): ", e, "\ng:", g)




}
