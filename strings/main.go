package main

import (
	"fmt"
	"strings"
)

func main() {
	string := "apple,banana,orange"
	a := strings.Split(string, ",")
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println(string)

	str := "one two nine two three five two"
	fmt.Println(strings.Count(str,"two"))


	str2:="                Rohan       Kumar       "
	fmt.Print(strings.TrimSpace(str2))
}
