package main

import (
	"fmt"
	"myLearning/myutils"
)

func main() {
	fmt.Print("Hello Rohan\n")
	myutils.PrintMessage("My message\n")

	// datatypes
	var name string = "Rohan"
	var age = 24
	println(name)
	println(age)

	var currency = 65000
	println(currency)

	// Shortcut variable define Syntex:
	person := "Rohan Kumar"
	println(person)
}
