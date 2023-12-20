package main

import "fmt"

const LoginToken string = "123456789"

func main() {
	var username string = "Preetam"
	fmt.Println(username)
	fmt.Printf("Variable is of the type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of the type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of the type: %T \n", smallVal)

	var smallFloat float32 = 255.123456789
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of the type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of the type: %T \n", anotherVariable)

	var anotherVariable2 string
	fmt.Println(anotherVariable2)
	fmt.Printf("Variable is of the type: %T \n", anotherVariable2)

	var anotherVariable3 bool
	fmt.Println(anotherVariable3)
	fmt.Printf("Variable is of the type: %T \n", anotherVariable3)

	// implicit type

	var website = "google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of the type: %T \n", website)

	// No var style
	numberOfUser := 25000
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of the type: %T \n", numberOfUser)

	fmt.Println("Login token is: ", LoginToken)
	fmt.Printf("Variable is of the type: %T \n", LoginToken)

}
