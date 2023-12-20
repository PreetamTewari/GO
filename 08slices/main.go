package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello World")

	// var fruits = []string{"apple", "banana", "grape"}

	// fmt.Println("fruits: ", fruits)
	// fmt.Printf("Type of fruits: %T\n", fruits)

	// fruits = append(fruits, "orange", "mango")
	// fmt.Println("fruits: ", fruits)

	// fruits = append(fruits[1:3])
	// fmt.Println("fruits: ", fruits)

	// fmt.Println("Length :", len(fruits))

	// fmt.Println("Capacity :", cap(fruits))

	// highScores := make([]int, 4)
	// highScores[0] = 234
	// highScores[1] = 945
	// highScores[2] = 465
	// highScores[3] = 867

	// fmt.Println("highScores: ", highScores)
	// highScores = append(highScores, 555, 666, 321)
	// fmt.Println("highScores: ", highScores)

	// sort.Ints(highScores)
	// fmt.Println(highScores)

	// fmt.Println(sort.IntsAreSorted(highScores))

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println("courses: ", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("courses: ", courses)
}
