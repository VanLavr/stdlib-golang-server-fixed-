package api

import (
	"crud/parser"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	id       int     `json:"ID"`
	Name     string  `json:"Name"`
	Category string  `json:"Category"`
	Price    float64 `json:"Price"`
}

func GetRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("|handled \"root\"| |request=GET| |status: %v|", http.StatusOK)

	page := parser.ParseFile("D:\\desktop2\\GoProjects\\stdlib-golang-server(fixed)\\frontend\\html\\root.html")
	fmt.Fprintf(w, page)
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	log.Printf("|handled \"hello\"| |request=GET| |status: %v|", http.StatusOK)

	page := parser.ParseFile("D:\\desktop2\\GoProjects\\stdlib-golang-server(fixed)\\frontend\\html\\index.html")
	fmt.Fprintf(w, page)
}

/*
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Printf("|handled \"all\"| |request=GET| |status: %v|", http.StatusOK)

}
*/
