package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["js"] = "JavaScript"
	languages["py"] = "Python"
	languages["go"] = "Golang"
	fmt.Println("languages:", languages)
	fmt.Println("js:", languages["js"])
	fmt.Println("py:", languages["py"])
	fmt.Println("go:", languages["go"])
	fmt.Println("len:", len(languages))

	delete(languages, "go")
	fmt.Println("languages: ", languages)

	for key, value := range languages {
		fmt.Printf("%s: %s\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("%s\n", value)
	}

}
