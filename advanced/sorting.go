package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

type sortingFunc func(p1, p2 *person) bool

type personSorter struct {
	people          []person
	sortingFunction sortingFunc
}

func (p *personSorter) Len() int {
	return len(p.people)
}

func (p *personSorter) Less(i, j int) bool {
	return p.sortingFunction(&p.people[i], &p.people[j])
}

func (p *personSorter) Swap(i, j int) {
	p.people[i], p.people[j] = p.people[j], p.people[i]
}

func (f sortingFunc) sorting(people []person) {
	ps := &personSorter{
		people:          people,
		sortingFunction: f,
	}
	sort.Sort(ps)
}

func main() {

	people := []person{
		{"Mohammad", 33},
		{"Reera", 25},
		{"Nazanin", 20},
		{"Mobina", 23},
	}

	sortingByAgeAscFunc := func(p1, p2 *person) bool {
		return p1.age < p2.age
	}

	sortingByNameAscFunc := func(p1, p2 *person) bool {
		return p1.name < p2.name
	}

	sortingByAgeDescFunc := func(p1, p2 *person) bool {
		return p1.age > p2.age
	}

	sortingByLengthOfNameAscFunc := func(p1, p2 *person) bool {
		return len(p1.name) < len(p2.name)
	}

	sortingFunc(sortingByAgeAscFunc).sorting(people)
	fmt.Println("Sorted by age:", people)

	sortingFunc(sortingByNameAscFunc).sorting(people)
	fmt.Println("Sorted by name:", people)

	sortingFunc(sortingByAgeDescFunc).sorting(people)
	fmt.Println("Sorted by age (descending):", people)

	sortingFunc(sortingByLengthOfNameAscFunc).sorting(people)
	fmt.Println("Sorted by length name:", people)

	names := []string{people[1].name, people[2].name, people[3].name, people[0].name}

	sort.Slice(names, func(i, j int) bool {
		return names[i][len(names[i])-1] < names[j][len(names[j])-1]
	})
	fmt.Println("Sorted by the last character of names:", names)

}

// Basic method to sort
// //Implementing the sort interface and its methods
// type person struct {
// 	name string
// 	age int
// }

// type peaopleList []person

// func (pl peaopleList) Len() int {
// 	return len(pl)
// }

// func (pl peaopleList) Less (i,j int) bool {
// 	return pl[i].age < pl[j].age
// }

// func (pl peaopleList) Swap (i,j int) {
// 	pl[i] , pl[j] = pl[j] , pl[i]
// }

// func main() {

// 	people := []person {
// 		{"Mohammad", 33},
// 		{"Reera", 25},
// 		{"Nazanin", 20},
// 		{"Mobina", 23},
// 	}

// 	sort.Sort(peaopleList(people))
// 	fmt.Println("Sorted by Age:", people)

// }

// func main(){

// 	numbers := []int{5,3,1,2,4}
// 	sort.Ints(numbers)
// 	fmt.Println("Sorted array: ", numbers)

// 	strings := []string{"Ali", "Mohammad", "Reera", "Nazanin", "Mobina"}
// 	sort.Strings(strings)
// 	fmt.Println("Sorted array: ", strings)

// }
