package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")
	myNumber := 123

	prt := &myNumber

	*prt += 2

	fmt.Println(prt, myNumber)
}