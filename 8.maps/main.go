package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps Modules")
	langauges := make(map[string]string)

	langauges["js"] = "JavaScript"
	langauges["py"] = "Python"
	langauges["go"] = "Go Lang"

	fmt.Println("Langauages", langauges)
	fmt.Println("js stands for : ", langauges["js"])

	delete(langauges, "go")

	fmt.Println("Updated Langauges", langauges)
}