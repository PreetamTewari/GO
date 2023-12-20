package main

import (
	"fmt"
	"time"
)

func main() {
	welcome := "Hello World"
	fmt.Println(welcome)

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.November, 24, 1, 1, 1, 1, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

}
