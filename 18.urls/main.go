package main

import (
	"fmt"
	"net/url"
)

const ipUrl = "https://ipinfo.io/2401:4900:1f24:8fff::369:13a5/geo"

func main() {
	fmt.Println("Welcome to url module")

	// parsing url
	result, err := url.Parse(ipUrl)


	if err != nil{
		panic(err)
	}

	fmt.Println("result of the url: ", result.Scheme)

	// construct a url
	googleUrlStruct := &url.URL{
		Scheme: "https",
		Host: "www.google.com",
	}

	googleUrl := googleUrlStruct.String()

	fmt.Println("google url is", googleUrl)

}