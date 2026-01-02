package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		m := make(map[string]int)

		m["hello"] = 123
		m["world"] = 456
		jsonAns, _ := json.Marshal(m)
		w.Write(jsonAns)
	}

	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}
