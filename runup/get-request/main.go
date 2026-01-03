package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := req.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	fmt.Fprintf(w, `{"message": "Hello, %s"}`, name)
}

func main() {
	http.HandleFunc("/hello", HelloHandler)
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
