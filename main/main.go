package main

import (
	"crud/DBconnection"
	"crud/api"
	"log"
	"net/http"
)

func main() {
	DBconnection.Connect()
	log.Printf("|listening and serving on localhost:8080|")

	// home endpoint
	http.HandleFunc("/", api.GetRoot)
	// hello endpoint
	http.HandleFunc("/hello", api.GetHello)
	// Get all info endpoint
	//http.HandleFunc("/all")

	serverError := http.ListenAndServe(":8080", nil)
	if serverError != nil {
		log.Fatal(serverError.Error())
	}
}
