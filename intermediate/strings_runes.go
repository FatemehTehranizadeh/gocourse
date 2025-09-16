package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/* Strings
	Sequence of bytes (unsigned int8 (uint8))
	immutable
	characters like \n are scappe sequences
	\n is for new line
	\t is for
	\r is for return to the first character (It's not used today because we have \n and with new line we go to the first of line)
	strings are made of rune characters
	String declaration: using double quotes
	String literals: using backticks

	*/
	/*
		Rune: each alphabet
		An onteger value that represents a character
		rune characters make strings
	*/
	message := "Hello \rGo!"    //prints "Go!sage: Hello "
	rawMessage := `Hello \nGo!` //backtick treats everything like a character

	fmt.Println("message:", message)
	fmt.Println("Raw message:", rawMessage)

	//length of a string
	fmt.Println("Length of message:", len(message))        //\r is counted as 1 - prints 10
	fmt.Println("Length of raw message:", len(rawMessage)) //\n is counted as two characters - prints 11

	//Indexing
	fmt.Println("The 4th character of message is:", message[3]) // Prints the ASCII code

	//Concatenating two strings
	greeting := "Hello "
	name := "Reera"
	fmt.Println(greeting + name) //Prints "Hello Reera"

	//Comparing two strings
	str1 := "Apple"  //ASCII value of A = 65
	str2 := "banana" //ASCII value of b = 98
	str3 := "App"
	str4 := "app"                                            //ASCII value of a = 97
	fmt.Println("Is Apple bigger than banana?", str1 > str2) //Prints false because of the alphabetical order in the dictionary => lexical
	// graphical comparison. Sorting, searching and other operations that involve comapring are used this method.
	fmt.Println("Is Apple bigger than App?", str1 > str3) //Prints true because of length
	fmt.Println("Is Apple bigger than app?", str1 > str4) //Prints false because the ASCII code of lowercase alphabets is higher than the upeercase ones

	//String iteration
	// for i, v := range message {
	// 	fmt.Printf("Index: %d , Value: %c\n", i, v)
	// }

	/*Formatting verbs:
	%d: for decimals, integers
	%c: for character values
	%x: hexadecimal format
	%v: general form - default value - The ASCII value of character - uint8 value
	%T: type of the variable
	*/

	//Counting the number of runes in a string
	fmt.Println("The number of runes in message is:", utf8.RuneCountInString(message))

	//Manipulating strings
	greetingWithName := greeting + name
	fmt.Println(greetingWithName)
	// We don't have append for strings

	/* Rune
	An alias for int32
	It's not a character, It's an integer value => A unicode code point
	Rune is declared with the type "Rune" and uses single quote

	*/

	var ch rune = 'a'
	fmt.Println("The ch value is:", ch) //Prints 97
	fmt.Printf("The ch real value is: %c\n", ch)

	jch := '日'
	fmt.Println("The jch value is:", jch) //Prints 26085
	fmt.Printf("The jch real value is: %c\n", jch)

	//Converting the runes to strings
	cstr := string(ch)
	fmt.Println("Converting rune to string:", cstr) //Prints a
	fmt.Printf("The type of ch is %T\n", ch)        //Prints int32
	fmt.Printf("The type of cstr is %T\n", cstr)    //Prints string

	smiley := '😊'
	fmt.Printf("The real value of smiley is %c and its rune value is %v\n", smiley, smiley)

}
