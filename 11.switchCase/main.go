package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to Switch Case Module")
	// rand.Seed(time.Now().UnixNano())
	diceVal := rand.Intn(6) + 1

	fmt.Printf("Dice Rolled And Value is %v \n", diceVal)

	switch diceVal {
	case 1:
		fmt.Println("You can either move out (OPEN) or move one step")
	case 2:
		fmt.Println("Move two steps")
	case 3:
		fmt.Println("Move three steps")
		fallthrough
	case 4:
		fmt.Println("Move four steps")
		fallthrough
	case 5:
		fmt.Println("Move five steps")
	case 6:
		fmt.Println("Move six steps and you can roll the DICE AGAIN !!")
	default:
		fmt.Println("Unlucky Please roll again")
	}
}