package main

import "fmt"

func main() {
	fmt.Println("hello structs")

	Aditya := User{"Aditya", "adi@gmail.com", true, 20}

	fmt.Println(Aditya)
	fmt.Printf("Aditya details are: %+v\n", Aditya)
	fmt.Printf("Aditya name is: %v\n", Aditya.Name)
}

type User struct {
	Name  string
	Email string
	Login bool
	Age   int
}
