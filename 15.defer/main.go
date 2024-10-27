package main

import "fmt"

func main() {
	fmt.Println("Welcome to Defer Module ")
	defer fmt.Println("\nI am the first line to be defered")
	fmt.Println("line after defer ")
	pushToDeferStack()
}

func pushToDeferStack() {
	for index := range(5){
		defer fmt.Print(index)
	}
}