package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	defer fmt.Println(add(10, 20)) // This line is execute just before code ends.
	fmt.Println("First Line")
	fmt.Println("Second Line")
	fmt.Println("Third Line")
	defer fmt.Println("Forth Line")
	defer fmt.Println("Fif th Line")

}

// Note : The defer keyword is act as stack. When it gets two defer LIFO excecution happens.

// 2. Defer is generally used before the function call so taht it will clode the file when complete code is getting executed.