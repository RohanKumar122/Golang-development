package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
}

func main() {
	fmt.Println("CRUD loading...")
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
