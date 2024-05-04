package main

import "fmt"

func main() {
	age := 23
	name := "Rio"
	height := 5.4322
	fmt.Println("name: ", name, "age: ", age, "Height: ", height)

	// Format Specifier (printf)

	fmt.Printf("Name is %s age is %d Height is %.3f", name, age, height)
}
