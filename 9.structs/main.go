package main

import "fmt"

type Cars struct{
	Name string
	Manufacturers string
	Year string
	IsAvailable bool
}

func main() {
	fmt.Println("Welcome to Strcuts Modules")

	swift := Cars{"Swift", "Maruti Suzuki", "2000-01-01", true}
	xuv300 := Cars{"XUV 300", "Mahnidar", "2024-01-01", true}

	fmt.Printf("Availble Cars are : %+v \n", swift)
	fmt.Printf("Availble Cars are : %+v \n ", xuv300)
}