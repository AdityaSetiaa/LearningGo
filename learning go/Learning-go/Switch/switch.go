package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print("Hello Switch\n")

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(6) + 1
	fmt.Println("Random number is:", num)

	switch num {
	case 1:
		fmt.Println("You rolled a one!")
	case 2:
		fmt.Println("You rolled a two!")
	case 3:
		fmt.Println("You rolled a three!")
	case 4:
		fmt.Println("You rolled a four!")
	case 5:
		fmt.Println("You rolled a five!")
	case 6:
		fmt.Println("You rolled a six! ROLL AGAIN")
	}
}
