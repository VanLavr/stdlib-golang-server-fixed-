package main

import (
	"log"
	"crud/api"
	"net/http"
)

func main() {
	log.Printf("|listening and serving on localhost:8080|")

	// home endpoint
	http.HandleFunc("/", api.GetRoot)
	// hello endpoint
	http.HandleFunc("/hello", api.GetHello)

	serverError := http.ListenAndServe(":8080", nil)
	if serverError != nil {
		log.Fatal(serverError.Error())
	}
}