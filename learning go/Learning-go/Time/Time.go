package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Print("Time")

	presentTime := time.Now()
	fmt.Println("Current time is: ", presentTime)

	fmt.Println("Current time is: ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.October, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Created date is: ", createdDate.Format("01-02-2006 Monday"))
}
