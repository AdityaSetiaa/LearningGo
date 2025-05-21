package main

import "fmt"

func main() {
	fmt.Println("Hello Slice")

	var fruits = []string{"Apple", "Banana", "Mango", "Orange", "Grapes"}
	fmt.Printf("type is %T\n", fruits)
	fruits = append(fruits, "Pineapple") // adds
	fruits = append(fruits[2:])
	fmt.Println(fruits)          // [Mango Orange Grapes Pineapple]
	fruits = append(fruits[1:3]) // [Orange Grapes]
	fmt.Println(fruits)          // [Orange Grapes]
	fruits = append(fruits[1:])  // [Grapes]

	highscore := make([]int, 4)

	highscore[0] = 100
	highscore[1] = 200
	highscore[2] = 300
	highscore[3] = 400
	highscore = append(highscore, 500, 600, 700)

}
