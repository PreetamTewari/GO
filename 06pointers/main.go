package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	var pointer = &myNumber
	fmt.Println("Value of pointer is ", pointer)
	fmt.Println("Value of pointer is ", *pointer)
	*pointer = *pointer + 2
	fmt.Println("New value is: ", myNumber)
}
