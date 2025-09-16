package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string  `json:"first_name"`
	Age       int     `json:"age,omitempty"`
	Email     string  `json:"-"`
	Address   Address `json:"address"`
}

type Address struct {
	City   string `json:"city"`
	Street string `json:"street,omitempty"`
	Number int    `json:"number"`
}

type Employee struct {
	FullName string `json:"full_name"`
	EmpID    string `json:"emp_id"`
	Age      int    `json:"age"`
	Address  `json:"address"`
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	p1 := Person{
		FirstName: "Reera",
		Age:       25,
	}

	jsonP1, err := json.Marshal(p1)
	checkError(err)
	fmt.Println(string(jsonP1))

	p2 := Person{
		FirstName: "Mohammad",
		Email:     "Mohammad@gmail.com",
		Address: Address{
			City:   "Tehran",
			Number: 3,
		},
	}
	jsonP2, err := json.Marshal(p2)
	checkError(err)
	fmt.Println("jsonP2 is:", string(jsonP2))

	jsonData1 := `{"full_name": "Jenny Dow", "emp_id":"0009", "age":30, "address":{"city":"San Jose", "Number": 1}}`
	var structFromJson Employee
	err = json.Unmarshal([]byte(jsonData1), &structFromJson)
	checkError(err)
	fmt.Println(structFromJson)
	fmt.Println(structFromJson.City)

	listOfCityAndStreets := []Address{
		{City: "Tehran", Street: "Ashrafi"},
		{City: "Kerman", Street: "Shafa"},
		{City: "Bandar-Abbas", Street: "Shalamcheh"},
		{City: "Tabriz", Street: "Abresan"},
	}
	jsonList, err := json.Marshal(listOfCityAndStreets)
	checkError(err)
	fmt.Println(string(jsonList))

	// handling unknown JSON structures
	var data map[string]interface{}
	jsonData2 := `{"first_name": "Jenny", "last_name": "Dow", "age":30, "address":{"city":"San Jose", "state": "CA"}, "isFemale": true}`

	err = json.Unmarshal([]byte(jsonData2), &data)
	checkError(err)

	fmt.Println(data)
	fmt.Println(data["last_name"])

}
