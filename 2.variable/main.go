package main

import (
	"fmt"
	"strings"
)

const Token = "JWT 1234567890"

func main() {
	var stringVal = "this is a String"
	intVal := 1234567
	floaVal := 3.1428

	fmt.Printf("%T \n, %T \n, %T \n, %T \n", stringVal, intVal, floaVal, Token)

	fmt.Println(strings.ContainsAny(stringVal, "z y x is"))
}
