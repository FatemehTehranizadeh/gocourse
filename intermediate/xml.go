package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	// City    string   `xml:"city,omitempty"`
	Email   string  `xml:"email"`
	Address Address `xml:"address"`
}

type Address struct {
	City   string `xml:"city"`
	Street string `xml:"street,omitempty"`
	Number int    `xml:"number"`
}

type Book struct {
	XMLName xml.Name `xml:"book"`
	ISBN    string   `xml:"isbn,attr"`
	Title   string   `xml:"title,attr"`
	Author  string   `xml:"author,attr"`
}

func main() {
	p1 := Person{
		Name: "Reera",
		Age:  26,
		// City: "Tehran",
		Email: "Reera@gmail.com",
		Address: Address{
			City:   "Tehran",
			Number: 3,
		},
	}
	xmlP1, err := xml.MarshalIndent(p1, "", "  ")
	if err != nil {
		log.Fatalln("Error while marshalling:", err)
	}
	fmt.Println(string(xmlP1))

	var p2 Person
	xmlRaw := `<person><name>John</name><age>30</age><address><city>Los Angeles</city><number>10</number></address></person>`
	// xmlRaw := `<person><name>John</name><age>30</age><city>Los Angeles</city><number>10</number></person>`

	err = xml.Unmarshal([]byte(xmlRaw), &p2)
	if err != nil {
		log.Fatalln("Error while unmarshalling:", err)
	}
	fmt.Println(p2)

	b1 := Book {
		ISBN: "46-598-564-23-26",
		Title: "Emma",
		Author: "Jane Austen",
	}

	xmlB1, err := xml.MarshalIndent(b1, "","  ")
	if err != nil {
		log.Fatalln("Error while marshalling:", err)
	}
	fmt.Println(p2)
	fmt.Println(string(xmlB1))

}
