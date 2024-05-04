package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter you name please: ")
	// // Method 1:
	// var name string
	// fmt.Scan(&name)
	// fmt.Printf("Hello mr. %s",name)

	// Method 2:
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s",name)
}
