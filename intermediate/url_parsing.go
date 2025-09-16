package main

import (
	"fmt"
	"net/url"
)


func main(){

	// Extracting information from URLs and APIs
	// [scheme://][userinfo@]host[:port][/path][?query][#fragment]

	rawURL1 := "https://example.com:8080/path/?query=param#fragment"
	// rawURL2 := "http://subdomain.example-long-domain.com:3000/path/to/page?search=go&lang=en"

	parsedURL, err := url.Parse(rawURL1)
	if err != nil {
		fmt.Println("Error in parsing the url: ", err)
	}
	fmt.Printf("The parsed URL is \n %v \nand its type is %T\n", parsedURL, parsedURL)

	fmt.Println("Scheme:" ,parsedURL.Scheme)
	fmt.Println("Host:" ,parsedURL.Host)
	fmt.Println("Port:" ,parsedURL.Port())
	fmt.Println("Path:" ,parsedURL.Path)
	fmt.Println("Query:" ,parsedURL.RawQuery)
	fmt.Println("Fragment:" ,parsedURL.Fragment)

	rawURL2 := "https://example.com/path?name=John&age=30"
	parsedURL2, err := url.Parse(rawURL2)
	if err != nil {
		fmt.Println("Error in parsing the url: ", err)
	}
	url2QueryParams := parsedURL2.Query()
	fmt.Println("The Query Parameters of URL2 is: ", url2QueryParams)
	fmt.Println("The value of key 'name' is:",parsedURL2.Query().Get("name"))
	fmt.Println("The value of key 'age' is:",url2QueryParams.Get("age"))

	baseURL1 := &url.URL{
		Scheme: "https",
		Host: "example.com:8080",
		Path: "/path",	
	}
	baseURL1Query := baseURL1.Query() // Initiate a value as an instance of Query()
	baseURL1Query.Set("name", "Reera") // Set query parameters
	baseURL1Query.Set("job", "engineer")
	baseURL1Query.Set("age", "25")
	baseURL1.RawQuery = baseURL1Query.Encode() // Encode the RawQuery field

	fmt.Println(baseURL1.String())

	// Another way to set the Query parameters

	baseURL2Values := url.Values{}

	baseURL2Values.Add("name", "Looli")
	baseURL2Values.Add("age", "5")
	baseURL2Values.Add("city", "Tehran")

	baseURL2Queries := baseURL2Values.Encode()

	baseURL2 := "https://eaxmple.com/search" + "?" + baseURL2Queries
	fmt.Println(baseURL2)

}