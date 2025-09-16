package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	str := "Hello Go!"

	//Length of string
	fmt.Println("The length of string is:", len(str))

	//Concatenation
	str1 := "Hello"
	str2 := "World!"
	result := str1 + " " + str2
	fmt.Println(result)

	//Indexing
	fmt.Println("From fisr character to the 7th character:", str[:7])

	// String conversion
	num := 18
	str3 := strconv.Itoa(num) //Itoa is used for converting integer to string
	fmt.Printf("The value of num is %v and its type is %T\n", str3, str3)

	// String spilitting
	// fruits := "apple, orange, banana, pineapple"
	fruits1 := "apple- orange- banana- pineapple"
	fruits_array := strings.Split(fruits1, "-")
	fmt.Println(fruits_array)

	// String joining
	countries := []string{"UAE", "Iran", "Iraq", "Qatar"}
	joined_countries := strings.Join(countries, "- ")
	fmt.Println(joined_countries) // Prints UAE- Iran- Iraq- Qatar

	// Checking for a certain string
	fmt.Println(strings.Contains(str, "llo "))

	// Replacing a string
	str4 := "Go HeGollo Go Go Go "
	replaced_string := strings.Replace(str4, "Go", "Wooooorlddddd!", -1)
	fmt.Println(replaced_string)

	str5 := "           HeLLO EVEryonE 054648786"
	fmt.Println(strings.TrimSpace(str5)) // It deletes the spaces at the beginning
	fmt.Println(strings.ToLower(str5))   // Prints            hello everyone
	fmt.Println(strings.ToUpper(str5))   // Prints            HELLO EVERYONE

	// Repeating a string
	str6 := strings.Repeat("foo ", 4)
	fmt.Println(str6) // Prints foo foo foo foo 

	// Counting the occurance of a string
	fmt.Println(strings.Count(str6, "o")) // Prints 8

	// Prefix
	fmt.Println(strings.HasPrefix(str,"he")) // false
	fmt.Println(strings.HasPrefix(str,"Hell")) // true
	fmt.Println(strings.HasPrefix(str,"He")) // true

	// Suffix
	fmt.Println(strings.HasSuffix(str,"o!")) // true


	// Regular Expression

	str7 := "He5llo, 123 Go! 11"
	re := regexp.MustCompile(`\d+`) // d means digits and + means digits from 1 to more
	reSlice := re.FindAllString(str7, -1) // -1 means all the matchs
	fmt.Println(reSlice) // [5 123 11]

	str8 := "Hello, ریرا"
	fmt.Println(utf8.RuneCountInString(str8)) //Prints 11

	// Building efficient strings using strings.Builder, It minimizes memory allocations and copies

	var builder strings.Builder
	
	// Write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")

	// Making the final string
	builderResult := builder.String()
	fmt.Println("The result of builder is:", builderResult)

	// Using writerrune to add a character
	builder.WriteRune('ر')
	builder.WriteString("How are you?")
	builderResult = builder.String()
	fmt.Println(builderResult)

	// Another Builder
	var builder1 strings.Builder

	builder1.WriteString("سلام")
	builder1.WriteRune(' ')
	builder1.WriteRune(' ')
	builder1.WriteRune(' ')
	builder1.WriteRune(' ')
	builder1.WriteString("ریرا")

	fmt.Println(builder1.String())

	builder1.Reset()

	builder1.WriteString("تکرار")
	builder1.WriteRune(' ')
	builder1.WriteRune(' ')
	builder1.WriteRune(' ')
	builder1.WriteRune(' ')
	builder1.WriteString("غریبانه")

	fmt.Println(builder1.String())







}
