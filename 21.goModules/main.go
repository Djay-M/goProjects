package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to go modules")

	rounter := mux.NewRouter()
	rounter.HandleFunc("/", getServerStatus).Methods("GET")
	rounter.HandleFunc("/api/status", getServerStatus).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", rounter))
}

func getServerStatus(response http.ResponseWriter, req *http.Request) {
	response.Write([]byte("GO Server is up and running !!"))
}