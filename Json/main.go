package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Is_adult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("This is JSON")
	person_data := Person{Name: "John", Age: 12, Is_adult: true}
	fmt.Println("Person is :", person_data)

	// convert person into json Encoding (Marhsalling)
	jsonData, err := json.Marshal(person_data)
	if err !=nil{
		fmt.Println(err)
        return
	}
	fmt.Println("Person is :",(jsonData))
	fmt.Println("Person is :", string(jsonData))

	// Decoding (Unmarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &person_data)
	if err!= nil {
        fmt.Println(err)
        return
    }
	fmt.Println("Person Data is: ",personData)

	// Type
	fmt.Printf("Type of person_data: %T\n", person_data)
	fmt.Printf("Type of jsonData: %T", jsonData)




}
