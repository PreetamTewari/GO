package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	loginCount := 20
	var result string
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println("Result:", result)

	if num := 2; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is greater than 10")
	}

}
