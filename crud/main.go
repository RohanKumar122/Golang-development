package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"complete"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Something went wring", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println(res.Status)
		return
	}

	// Method 1: It return JSON then we have to decode using marshal
	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Something went wrong", err)
	// 	return
	// }
	// fmt.Println(string(data))

	// Method2:
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error Decoding: ", err)
		return
	}
	fmt.Println("Todo", todo)
}

func performPostRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "Rohan Kumar",
		Completed: true,
	}

	// Convert the todo struct to JSON
	// res, err = http.Post("https://jsonplaceholder.typicode.com/todos/")

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Erro Marshaling ", err)
		return
	}
	// Convert Json data to string
	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	// POST REQUEST
	res, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error in Post request", err)
		return
	}
	defer res.Body.Close()

	// convert response in readable format (As POST request give response)
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in POST response", err)
		return
	}
	fmt.Println("Response is: ", string(data))
	fmt.Println("Response Status: ", res.Status)

}

func main() {
	fmt.Println("CRUD loading...")
	// performGetRequest()
	performPostRequest()

}
