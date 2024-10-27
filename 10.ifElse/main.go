package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to if-else Module in GO")
	var nums = rand.Perm(100001)
	// sort.Ints(nums)
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a number, to check if it in the nums array")

	inputNum, _ := reader.ReadString('\n')
	number, _ := strconv.ParseInt(strings.TrimSpace(inputNum), 10, 64)
	isFound := false

	fmt.Println("Your Number was", number)
	for val := range nums {
		if val == int(number) {
			isFound = true
			break
		}
	}
	if isFound {
		fmt.Println("Found Your Number, You win the Lottery")
	} else {
		fmt.Println("Sorry the number was not found, please try again in few mins")
	}
}