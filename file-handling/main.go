package main

import (
	"fmt"
	"os"
)

func main() {

	// --------- write file --------

	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("Error while creating file")
	// 	return
	// }
	// defer file.Close()
	// content := "Hello Rohan!"
	// // file.Write([]byte(content))  //Method 1 (self)

	// // Method 2
	// _, errors := io.WriteString(file, content+"\n")
	// if errors != nil {
	// 	fmt.Println("Error while writing file: ", errors)
	// 	return
	// }

	// fmt.Println("Successfully Created file!!")



	// -------- READ file--------
	// there are basically two methods to read the file : a. load complete file in memory and  b. load in chuncks using buffer.

	// METHOD 1: load complete file
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Erro while reading the file.", err)
		return
	}
	fmt.Println(string(content))

}
