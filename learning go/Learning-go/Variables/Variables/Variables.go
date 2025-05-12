package main

import "fmt"

const Loggedin bool = true

func main() {
	//----------------------------------------Learining go variables------------------------------------
	fmt.Println("Hello world")

	var username string = "Aditya"
	fmt.Println(username)
	fmt.Printf("variable is : %T \n", username)

	var isLogedin bool = true
	fmt.Println(isLogedin)
	fmt.Printf("variable is : %T \n", isLogedin)

	var smallFloat float32 = 259.4578907654
	fmt.Println(smallFloat)
	fmt.Printf("variable is : %T \n", smallFloat)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("variable is : %T \n", smallval)

	var variable int
	fmt.Println(variable)
	fmt.Printf("variable is : %T \n", variable)

	myname := "Aditya"
	fmt.Println(myname)
	fmt.Printf("variable is : %T \n", myname)

	fmt.Printf("variable is : %T \n", Loggedin)
	fmt.Println(Loggedin)
}
