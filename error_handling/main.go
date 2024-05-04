package main

import "fmt"

func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return (a / b), nil
}
func main() {
	// fmt.Println("This is Error handling.")
	// ans, err:= divide(10,1)

	// if(err!=nil){
	// 	fmt.Println("Error handling")
	// }

	// fmt.Println(ans)

	// OR

	ans, _ := divide(10, 0)
	fmt.Println(ans)
}
