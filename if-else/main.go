package main

import (
	"fmt"
)

func main() {

	var a= 5
	
	if a==5{
		fmt.Println("Number is Equal to 5")
	}else if a<5 {
		fmt.Println(a," is less than 5")
	}else{
		fmt.Println(a," is greater than 5")
	}
}