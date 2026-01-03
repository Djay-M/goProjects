package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup
var websitesVisited []string
var mux sync.Mutex

func main() {
	startTime := time.Now()
	fmt.Println("Welcome to GO Routies Module")
	// go printMsg("Hello")
	// printMsg("World")

	websites := []string{
		"https://google.com",
		"https://fb.com",
		"https://loc.dev",
		"https://github.com",
		"https://djay-m.github.io/resume/",
		"http://localhost:4000/api/v1/status",
	}

	for _, website := range websites {
		go getWebStatus(website)
		waitGroup.Add(1)
	}

	waitGroup.Wait()

	endTime := time.Now()
	fmt.Println(websitesVisited)
	fmt.Println(endTime.Sub(startTime))
}

// func printMsg(s string) {
// 	for range 5 {
// 		time.Sleep(5 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getWebStatus(url string) {
	defer waitGroup.Done()
	result, err := http.Get(url)

	if err != nil {
		log.Fatal("Error while calling the URL", url)
	}

	mux.Lock()
	websitesVisited = append(websitesVisited, url)
	mux.Unlock()
	fmt.Printf("%d is status code for %s \n", result.StatusCode, url)
}
