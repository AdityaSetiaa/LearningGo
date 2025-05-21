package main

import "fmt"

func main() {
	fmt.Println("Hello Pointers")

	var ptr *int
	fmt.Println(ptr) //nil

	mynum := 23
	var ptr1 *int = &mynum
	fmt.Println(ptr1)  //0xc0000b2008
	fmt.Println(*ptr1) //23
}
