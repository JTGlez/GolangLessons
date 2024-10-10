package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type GreetingRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("pong"))
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
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", pingHandler)

	r.Route("/greetings", func(r chi.Router) {
		r.Post("/", greetingsHandler)
	})

	fmt.Println("Server listening at: localhost:8081")
	if err := http.ListenAndServe(":8081", r); err != nil {
		fmt.Println("Error while starting server:", err)
	}
}
