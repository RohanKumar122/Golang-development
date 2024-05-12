package main

import (
	"fmt"
	"time"
)

func First(){
	fmt.Println("This is First")
}

func Second() {
	fmt.Println("This is Second")
	time.Sleep(2000* time.Millisecond)
}

func Third(){
	fmt.Println("This is Third")
}

func main() {
	First()
	go Second()
	Third()
	time.Sleep(1000* time.Millisecond)

}