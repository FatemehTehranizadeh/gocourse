package main

import (
	"fmt"
	"regexp"
)

func main() {

	// re := regexp.MustCompile(`[a-zA-z0-9._%-]+@[a-zA-z0-9.-]+\.[a-zA-z]{2,}`)
	// re := regexp.MustCompile(`[a-zA-z\d._%-]+@[a-zA-z\d.-]+\.[a-zA-z]{2,}`)

	// email1 := "fkddklf5566@gmail.uk.com"
	// email2 := "gdklsdcjs;dl.fdfk"

	// fmt.Println("Email1:", re.MatchString(email1))
	// fmt.Println("Email2:", re.MatchString(email2))

	// date regex
	// dateRE := regexp.MustCompile(`\d{4}/\d{1,2}/\d{1,2}`)
	// dateRE := regexp.MustCompile(`\d{4}/(0[1-9]|1[0-2])/(0[1-9]|[12][0-9]|3[01])`)
	dateRE := regexp.MustCompile(`(13[0-9]{2}|14[0-4]{2})/(0?[1-9]|1[0-2])/(0?[1-9]|[12][0-9]|3[0-1])`)
	fmt.Println(dateRE.MatchString("1304/8/30"))

	exampleString := "Hello World!"
	vowelsRE := regexp.MustCompile("[aeoui]")
	res := vowelsRE.ReplaceAllString(exampleString,"8")
	fmt.Println(res)

	// flags
	// i => case insensitive
	insensitiveRE := regexp.MustCompile("(?i)go")
	exampleString = "Golang GOGOGOGOGOGOG"
	fmt.Println(insensitiveRE.MatchString(exampleString))

}
