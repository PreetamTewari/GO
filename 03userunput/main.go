package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the playground!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the course: ")

	// comma ok || comma error
	input, _ := reader.ReadString(('\n'))
	fmt.Println("Thanks for the rating: ", input)
	fmt.Printf("Type of rating is: %T \n", input)
}
