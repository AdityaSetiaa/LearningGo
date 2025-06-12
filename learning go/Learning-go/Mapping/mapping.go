package main

import "fmt"

func main() {
	fmt.Println("Hello Mapping")

	Languages := make(map[string]string)

	Languages["JS"] = "JavaScript"
	Languages["PY"] = "Python"
	Languages["RB"] = "Ruby"
	Languages["GO"] = "GoLang"
	fmt.Println("Languages is: ", Languages)
	fmt.Println("JS short form is: ", Languages["JS"])
	delete(Languages, "RB")

}
