package main

import "fmt"

func main() {
	day := 2

	switch day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Default output")

	}

	month :="January"
	switch month{
	case "January", "Fabruary", "March":
		fmt.Println("Spring")

	case "April","May","June":
		fmt.Println("Spring")

	case "July" ,"August","Septmber":
		fmt.Println("Summer")

	default:
		fmt.Println("Other Months")
	}
	
}