package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/json1", respondJson1)
	http.HandleFunc("/json2", respondJson2)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func respondJson1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Hello World!"}`)
}

type Message struct {
	Text    string
	IsValid bool
}

func respondJson2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := Message{"Hello World!", true}

	str, err := json.MarshalIndent(message, "", "    ")
	fmt.Fprintf(w, string(str))
	if err != nil {
		fmt.Fprintf(w, "Error encountered")
	}
}
