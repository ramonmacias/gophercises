package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ramonmacias/gophercises/cyoa/handler"
)

func main() {
	m := mux.NewRouter()
	m.HandleFunc("/chapter/{id}", handler.ChapterHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", m)
}
