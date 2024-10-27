package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops Module")

	language := []string{"c++", "Java", "Python", "Node JS", "GO Lang"}
	language = append(language, "React JS")
	initalVal := 0

	// for index:= 0; index < len(language); index++ {
	// 	fmt.Println(language[index])
	// }

	// for i := range(language){
	// 	fmt.Println(language[i], i)
	// }

	// for _, lang := range(language){
	// 	fmt.Println(lang)
	// }

	fmt.Println(len(language))
	for initalVal < len(language) {
		fmt.Println(language[initalVal])
		initalVal++
	}
}