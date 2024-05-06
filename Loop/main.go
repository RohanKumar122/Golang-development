package main

import "fmt"

func main() {
	count := 5
	for i := 0; i <= count; i++ {
		fmt.Println(i)
	}

	// Infinite Loop
	counts := 0
	for {
		fmt.Println("Infinite Loop")
		if counts == 5 {
			break
		}
		counts += 1
	}

	//  RaNGE IN lOOP
	numbers :=[]int {11,22,23,4,5}
	for index,value := range numbers{
		fmt.Println(index,value*5)
	}

	data := "Hello, world!"

	for index,char:= range data{
		fmt.Printf("Index of Data is: %d and Character is: %c\n",index,char)
	}

	
}
