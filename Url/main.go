package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	myURL := "https://fakestoreapi.com/products/1"

	parseURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Error parsing", err)
		return
	}
	fmt.Println(parseURL)
	data, err := http.Get(myURL)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)

	// fmt.Println(string(data))  // This is not working that's why we use the below method ; ioutil.ReadAll

	// data1,err := ioutil.ReadAll(data.Body)
	// if err!= nil {
	//     fmt.Println(err)
	//     return
	// }

	// fmt.Println(string(data1))


	fmt.Printf("Type of URL: %T\n",data)
	fmt.Println("Scheme: ",parseURL.Scheme)
	fmt.Println("Host: ",parseURL.Host)
	fmt.Println("Path: ",parseURL.Path)

	//modify
	parseURL.Scheme = "http"
	parseURL.Host = "www.google.com"
    fmt.Println("Modified Scheme: ",parseURL.Scheme)
    fmt.Println("Modified Host: ",parseURL.Host)
    fmt.Println("Modified Path: ",parseURL.Path)
    fmt.Println("Modified URL: ",parseURL.String())

	//constructing a URL string from URL object
	newURL :=parseURL.String()

	fmt.Println("New Url: ",newURL)
	

}
