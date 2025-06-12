package main

import "fmt"

// Structs are user defined data types in Go
func main() {
	fmt.Println("Hello Loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("Index: %d, Day: %s\n", index, day)
	}

	rougueValue := 1

	for rougueValue < 10 {
		fmt.Println("Rogue value is:", rougueValue)
		rougueValue++
	}
}
