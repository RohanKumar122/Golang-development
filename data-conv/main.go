package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num = 42
	fmt.Printf("Number is %d\n", num)
	fmt.Printf("Type is %T\n", num)

	var data float64 = float64(num)
	fmt.Printf("Type is %T\n", data)

	num = 123
	str := strconv.Itoa(num)
	fmt.Printf("Number is %s\n", str)
	fmt.Printf("Number is %T\n", str)

	number_string :="1234"
	number_int,_ := strconv.Atoi(number_string)

	fmt.Printf("Number is %d\n", number_int)
	fmt.Printf("Number is %T\n", number_int)

}
