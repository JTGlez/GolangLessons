package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GreetingRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "pong")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func greetingsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Method: %s\n", r.Method)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request GreetingRequest
	body, err := io.ReadAll(r.Body)
	fmt.Printf(string(body))
	if err != nil {
		http.Error(w, "Error while reading body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Error while decoding json", http.StatusBadRequest)
		return
	}

	fmt.Printf("Request: %+v\n", request)

	response := fmt.Sprintf("Hello %s %s", request.FirstName, request.LastName)
	fmt.Fprint(w, response)
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/greetings", greetingsHandler)

	fmt.Println("Server listening at: localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Errow while :", err)
	}
}
