package main

import "fmt"

func main() {
	fmt.Println("This is Slices")

	number := []int{1, 2, 3, 4, 5}

	fmt.Printf("Number is datatype of %T\n", number)
	fmt.Printf("Lenght of datatype: %T\n", len(number))

	name := []string{"Rohan"}
	fmt.Printf("Name is: %s", name[0])

	// Make function
	numbers :=make([] int,3,5)
	numbers = append(numbers,4)
	fmt.Println("Slices: ",numbers)
	fmt.Println("Length", len(numbers))
	fmt.Println("Capacity: ",cap(numbers))





}
