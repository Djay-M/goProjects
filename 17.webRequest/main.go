package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const googleUrl = "https://ipinfo.io/2401:4900:1f24:8fff::369:13a5/geo"//"https://www.google.com/"

func checkForError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Welcome to web request module")
	response, webErr := http.Get(googleUrl)

	checkForError(webErr)

	responseBytes, ioErr := ioutil.ReadAll(response.Body)

	checkForError(ioErr)

	responseBody := string(responseBytes)
	fmt.Println("response recevied from google url is", responseBody)
}