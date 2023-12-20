package main

import "fmt"

func main() {
	preetam := User{"Preetam", "preetam@gmail.com", true, 32}
	fmt.Println("User:", preetam)
	fmt.Printf("User: %+v\n", preetam)
	fmt.Printf("Name is %v and email is %v\n", preetam.Name, preetam.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
