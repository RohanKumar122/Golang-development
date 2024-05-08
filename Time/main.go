package main

import (
	"fmt"
	"time"
)

func main() {
	current_time := time.Now()
	fmt.Println(current_time)
	fmt.Println("Hello, world!")
	fmt.Printf("Type of Time is %T\n", current_time)

	formated := current_time.Format("02-01-2006 15:04:05")
	formated1 := current_time.Format("02/01/2006 15:04:05")
	fmt.Println(formated)
	fmt.Println(formated1)

	layout_str :="02/01/2006"
	dateStr :="25/11/2030"
	formated_time,_ :=time.Parse(layout_str,dateStr)
	fmt.Println("Formatted time",formated_time)
}
