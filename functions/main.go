package main

import "fmt"

func Simplefunction() {
	fmt.Print("This is Simple Function.\n")
}

// func add(a,  b int )int{      //Add int type at the last position
// 	return a + b
// }
// OR
func add(a,  b int )(result int){      //Add int type at the last position
	result= a + b
	return 
}

func multiply(a,b int) int{
	return a*b
}

func main() {
	fmt.Println("This is function in Golang")
	Simplefunction()

	ans :=add(3,4)
	fmt.Printf("Sum of the numbers is %d\n",ans)

	data := multiply(3,4)
	fmt.Printf("Multiplication of the numbers is %d",data)
	
}
