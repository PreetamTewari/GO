package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello World")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice Number:", diceNumber)

	if diceNumber == 6 {
		fmt.Println("You Won!")
	} else {
		fmt.Println("You Lost")
	}

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
		fallthrough
	case 3:
		fmt.Println("You can move 3 spot")
		fallthrough
	case 4:
		fmt.Println("You can move 4 spot")
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot and roll dice again")
	default:
		fmt.Println("What was that!")
	}

}