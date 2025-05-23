package main

import "fmt"


func main(){
	// var map1 map[string]int
	// map1 = make(map[string]int)

	map1 := make(map[string]int)

	map1["key1"] = 1
	map1["key2"] = 2
	map1["key3"] = 3
	
	fmt.Println(map1["key2"])
	delete(map1, "key3")
	fmt.Println(map1)
	
	map1["key4"] = 4
	map1["key5"] = 5
	map1["key6"] = 6
	
	// clear(map1)
	// fmt.Println(map1) //Printing a blank map
	
	keyToCheck := "key14"
	value, isThereAnyValue := map1[keyToCheck]
	if isThereAnyValue {
		fmt.Printf("\nIs there any value for key '%v'? %v and the value is %v\n",keyToCheck, isThereAnyValue, value)
		} else {
			fmt.Println("There is no such this key!!!")
		}
		
		map2 := map[string]int{"a": 1, "b":2}
		fmt.Println(map2)
		
		var map3 *map[string]int
		map3 = &map2
		map2 ["random"] = 6
		fmt.Println("map3:", *map3)
		
		for k, v := range map2 {
			fmt.Printf("The key is '%v' and the value is '%v'\n", k, v)
		}
		
		//Convention is to use "ok" keyword
		keyToCheck = "key1"
		value, ok := map1[keyToCheck]
		if ok {
			fmt.Println("The value exists with ",keyToCheck)
			}else {
				fmt.Println("The value does NOT exists with ",keyToCheck)
			}
		
	//Initialization value of maps is NIL
	var map4 map[string]int
	fmt.Println(map4)

	map5 := map[int]bool {0: false, 1:true, 2: false, 3:false}
	fmt.Println("map5:", map5)
	fmt.Println("map5[2]:",map5[2])
	fmt.Println("The length of map5 is: ", len(map5)) //The number of keys the map has
	
	//nested maps
	map6 := make(map[string]map[string]string)
	map7 := map[string]string {"key": "Val"}
	map6 ["map1"] = map7
	fmt.Println("Map6 is:", map6)
}