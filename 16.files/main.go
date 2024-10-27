package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const DatabaseFilePath = "./database.txt"
func checkNilError(err error) {
	if err != nil{
		panic(err)
	}
}

func main() {
	fmt.Println("Welcome to files modules")
	fmt.Println("Please enter a message to update in the database")
	reader := bufio.NewReader(os.Stdin)
	input, readerErr := reader.ReadString('\n')

	checkNilError(readerErr)

	file, fileErr := os.OpenFile(DatabaseFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	checkNilError(fileErr)

	defer file.Close()
	
	writeLength, writeErr := io.WriteString(file, input)

	checkNilError(writeErr)

	fmt.Println("The length of chars added in database file are", writeLength)

	readFile(DatabaseFilePath)
}

func readFile(fileName string) {
	dataBytes, readFileErr := os.ReadFile(fileName)
	checkNilError(readFileErr)

	fmt.Println("database file data: ", string(dataBytes))
}