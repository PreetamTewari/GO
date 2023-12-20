package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println("Day", d+1, "is", days[d])
	// }

	for i := range days {
		fmt.Println("Day", i+1, "is", days[i])
	}

	for _, day := range days {
		fmt.Println("Day", day)
	}

	name := "Salaar "
	loop := 100
	for d := 0; d < loop; d++ {
		fmt.Print(name)
	}

}
