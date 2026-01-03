package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices Module")
	var nums = rand.Perm(11)
	var removeVal = 5
	sort.Ints(nums)
	fmt.Println(nums)

	nums = append(nums[:removeVal], nums[removeVal+1:]...)
	fmt.Println(nums)
}