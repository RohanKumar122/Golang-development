package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error while creating file")
		return
	}
	defer file.Close()
	content := "Hello Rohan!"
	// file.Write([]byte(content))  //Method 1

	// Method 2
	_, errors := io.WriteString(file, content+"\n")
	if errors != nil {
		fmt.Println("Error while writing file: ", errors)
		return
	}

	fmt.Println("Successfully Created file!!")
}
