package main

import "fmt"

func main() {
	println("Hello, World!")

	var fruits [3]string
	fruits[0] = "Apple"
	fruits[2] = "Banana"

	fmt.Println("Fruits:", fruits)
	fmt.Println(len(fruits))

	var vegList = [3]string{"Potato", "Tomato", "Onion"}
	fmt.Println("Vegetables:", vegList)
	fmt.Println("Length:", len(vegList))

}
