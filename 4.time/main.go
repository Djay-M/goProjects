package main

import (
	"fmt"
	"time"
)
const TimeFormate = "01-02-2006 Monday 15:04:05"
func main() {
	const welcomeMsg = "Welcome To Time Module"
	fmt.Println(welcomeMsg)

	todayDate := time.Now()

	fmt.Println(todayDate.Format(TimeFormate))
}