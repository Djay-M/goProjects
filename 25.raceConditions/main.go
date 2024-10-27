package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to race conditions in GO Lang")
	wt := &sync.WaitGroup{}
	mux := &sync.Mutex{}

	prices := []int{}

	// function 1
	wt.Add(1)
	go func(wt *sync.WaitGroup) {
		fmt.Println("Printing From 1")

		mux.Lock()
		prices = append(prices, 1)
		mux.Unlock()

		wt.Done()
	}(wt)

	// function 2
	wt.Add(1)
	go func(wt *sync.WaitGroup) {
		fmt.Println("Printing From 2")

		mux.Lock()
		prices = append(prices, 2)
		mux.Unlock()

		wt.Done()
	}(wt)

	// function 3
	wt.Add(1)
	go func(wt *sync.WaitGroup) {
		fmt.Println("Printing From 3")

		mux.Lock()
		prices = append(prices, 3)
		mux.Unlock()

		wt.Done()
	}(wt)

	wt.Wait()
	fmt.Println(prices)
}
