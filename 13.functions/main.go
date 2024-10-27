package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to Function Module")
	fmt.Println("The total after addition is: ", adder(2, 5))
	fmt.Println("The total after adding all values is: ", slicAdder(rand.Perm(11)...))
}


func adder(val1 int, val2 int) int {
	return val1 + val2
}

func slicAdder(values ...int) int {
	total := 0

	for _, val := range(values){
		total += val
	}

	return total
}