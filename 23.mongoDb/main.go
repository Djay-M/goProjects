package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Djay-M/mongoDB/routers"
)

func main() {
	fmt.Println("Welcome to mongo db module")
	fmt.Println("Starting server ...")

	router := routers.Router()

	fmt.Println("Server is running on port 4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}
