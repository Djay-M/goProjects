package main

import (
	"encoding/json"
	"fmt"
)

func handleError(err error) {
	if err != nil{
		panic(err)
	}
}

type Users struct {
	Name string `json:"customerName"`
	Phone string `json:"phoneNumber"`
	Password string `json:"-"` // do not show when converted into a JSON
	Tags []string `json:"tags,omitempty"`
	IsArchived bool `json:"isArchived"`
}

func main() {
	fmt.Println("Welcome to json module")
	customers := encodeJson()
	fmt.Printf("customers %s \n ", customers)
	decodeJson()
}

func encodeJson() []byte {
	customers := []Users {
		{Name: "Dev test", Phone: "1234567890", Password: "devtest1234567890", Tags: []string {"Node developer", "dev env"}, IsArchived: true},
		{Name: "QA test", Phone: "1234567890", Password: "qatest1234567890", Tags: []string {"QA", "QA env"}, IsArchived: true},
		{Name: "UAT test", Phone: "1234567890", Password: "uattest1234567890", IsArchived: true},
	}

	customersJson, jsonErr := json.MarshalIndent(customers, "", "\t")
	handleError(jsonErr)

	// fmt.Printf("customers has Json %s\n ", customersJson)
	return customersJson
}

func decodeJson() {
	webDataBytes := []byte(`
	{
		"Name": "QA test",
		"Phone": "1234567890",
		"Password": "qatest1234567890",
		"Tags": ["QA", "QA env"],
		"IsArchived": true
	}
	`)
	isValidJson := json.Valid(webDataBytes)

	var webJsonUser map[string]interface{}

	if !isValidJson{
		fmt.Println("webDataBytes is not a valid JSON")
	}

	json.Unmarshal(webDataBytes, &webJsonUser)
	fmt.Printf("webJson data is: %#v\n", webJsonUser)
}