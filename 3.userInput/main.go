package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please provide your rating")
	
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	} else {
		rating, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
		fmt.Println("Thank You For Your Rating", rating + 1)
	}

}