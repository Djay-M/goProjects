package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const testUrl = "http://localhost:3090"

func handleError(err error) {
	if err != nil{
		panic(err)
	}
}

func main() {
	fmt.Println("Welcome to script for calling a get api")
	getDataFromGoogle()
	createUser()
}

func getDataFromGoogle() {
	getTestUrl := testUrl + "/status"
	response, requestErr := http.Get(getTestUrl)
	handleError(requestErr)
	var responseString strings.Builder
	defer response.Body.Close()

	fmt.Println("status code: ", response.StatusCode)
	fmt.Println("content length is: ", response.ContentLength)

	contentBytes, responseErr := io.ReadAll(response.Body)
	handleError(responseErr)
	contentByteCount, contentByteCountErr := responseString.Write(contentBytes)
	handleError(contentByteCountErr)

	fmt.Println("content of the response body is ", contentByteCount, responseString.String())
}

func createUser() {
	createUserUrl := testUrl + "/api/v1/users/create"
	var responseString strings.Builder
	requestBody := strings.NewReader(`
		{
		"firstName": "player 4",
		"lastName": "dev player",
		"username": "player4",
		"password": "player4"
		}
	`)

	response, createUserErr := http.Post(createUserUrl, "application/json", requestBody)
	handleError(createUserErr)
	defer response.Body.Close()

	responseBytes, responseConvErr := io.ReadAll(response.Body)
	handleError(responseConvErr)

	responseBytesCount, resresponseBytesErr := responseString.Write(responseBytes)
	handleError(resresponseBytesErr)

	fmt.Println("new user created with the following details", responseBytesCount, responseString.String())
}
