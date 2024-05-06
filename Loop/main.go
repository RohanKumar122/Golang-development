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
}
