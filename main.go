package main

import (
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/stores", GetStore)
	http.HandleFunc("/path", CalculatePath)
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func main() {
	handleRequests()
}
