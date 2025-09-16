package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type person struct {
	PersonName string
	Personage  int
}

func main() {

	/* Text templates
	Define and execute templates for generating text outputs
	They are useful for generating structured texts such as HTML, JSON, SQL Queries
	Template: A string or file that contains one or more action sequences
	Actions examples: inserting values, iterating over data, executing conditionals
	The actions are enclosed in double curly braces
	Variable insertion => dot notation => {{.fieldName}}
	Templates are executed by applying them to data structures like structs, slices, maps
	text/templates has basic features
	HTML/templates has advance methods
	*/

	// tmpl := template.New("exampleName")
	// tmpl, err := tmpl.Parse("Welcome, {{.personName}}! How are you doing?")
	tmpl, err := template.New("exampleName").Parse("Welcome, {{.PersonName}}! How are you doing?\n")

	if err != nil {
		panic(err)
	}

	p1 := person{
		PersonName: "Reera",
	}

	data := map[string]interface{}{
		"PersonName": "Kotlet",
	}

	exeErr := tmpl.Execute(os.Stdout, p1)
	if exeErr != nil {
		panic(exeErr)
	}

	exeErr = tmpl.Execute(os.Stdout, data)
	if exeErr != nil {
		panic(exeErr)
	}

	// We can use template.Must() to handle the error

	tmpl1 := template.Must(template.New("exampleName").Parse("See you later {{.PersonName}}!\n"))
	tmpl1.Execute(os.Stdout, data)

	// Creating a console program

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	//Email template
	emailTmpl := template.Must(template.New("emailTemplate").Parse("Your email is: {{.emailAddress}}@gmail.com\n"))
	fmt.Println("Enter your gmail:")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	emaildata := map[string]string{
		"emailAddress": email,
	}
	err = emailTmpl.Execute(os.Stdout, emaildata)


	// templates := map[string]string{
	// 	"welcome":      "Welcome {{.name}}! we are glad to have you",
	// 	"notification": "Dear {{.name}}, you have a notification: {{.notification}}",
	// 	"error":        "error: {{.error}}",
	// }

	// parsedTemplates := make(map[string]*template.Template)
	// for tmplName,tmplStrings := range templates {
	// 	parsedTemplates[tmplName] = template.Must(template.New(tmplName).Parse(tmplStrings))
	// }

	// for {
	// 	fmt.Println("\nMenu:")
	// 	fmt.Println("1. Join")
	// 	fmt.Println("2. Get Notification")
	// 	fmt.Println("3. Get Error")
	// 	fmt.Println("4. Exit")
	// 	choice, _ := reader.ReadString('\n')
	// 	choice = strings.TrimSpace(choice)

	// 	var data map[string]string
	// 	var finalTmpl *template.Template

	// 	switch choice {
	// 	case "1":
	// 		data = map[string]string {"name" : name}
	// 		finalTmpl = parsedTemplates["welcome"]

	// 	case "2":
	// 		fmt.Println("Enter your notification please: ")
	// 		notification, _ := reader.ReadString('\n')
	// 		notification = strings.TrimSpace(notification)
	// 		data = map[string]string {"name" : name, "notification" : notification}
	// 		finalTmpl = parsedTemplates["notification"]

	// 	case "3":
	// 		data = map[string]string {"error" : "This is a test error"}
	// 		finalTmpl = parsedTemplates["error"]

	// 	case "4":
	// 		return

	// 	default:
	// 		fmt.Println("Invalid Choice")
	// 		continue
	// 	}

	// 	err = finalTmpl.Execute(os.Stdout,data)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// }

}
