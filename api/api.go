package api

import (
	"io"
	"net/http"
	"log"
	"crud/parser"
	"fmt"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("|handled \"root\"| |request=GET| |status: %v|", http.StatusOK)
	io.WriteString(w, "root...\n")
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	log.Printf("|handled \"hello\"| |request=GET| |status: %v|", http.StatusOK)

	page := parser.ParseFile("C:\\Users\\Ivan\\GoProjs\\CRUD\\frontend\\html\\index.html")
	fmt.Fprintf(w, page)
}