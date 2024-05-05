package main

import "fmt"

func main() {
	fmt.Println("Learning Array")
	// Method 1: 
	var name[5] string

	// mathod 2:
	// var name =[5] int{1,2,3,4,5}


	name[0]="Rohan"
	name[1]="Akash"

	fmt.Printf("There name are %q\n",name)
	fmt.Printf("There name are %s\n",name)
	fmt.Println("Lenght of Array is: ",len(name))
}
