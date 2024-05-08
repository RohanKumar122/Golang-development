package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://fakestoreapi.com/products/1")
	if err != nil {
		fmt.Println("Something went wrong", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("Type of response  is: %T\n", res)
	fmt.Println(res.StatusCode)
	//fmt.Println(res)

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Something went wrong", err)
		return
	}
	fmt.Printf("Type of data  is: %T\n", string(data))
	fmt.Println("response: ", string(data))

}
