package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type requestBody struct {
	Task string `json:"task"`
}

var task string

func createHandler(w http.ResponseWriter, r *http.Request) {
	var rq requestBody

	if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
		fmt.Printf("%v", "Error not decode from Body to reqBody struct")
		return
	}

	fmt.Println("successful decode")

	task = rq.Task

	fmt.Println("json data write in the task variable")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func getHandlers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"hello": task})

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", createHandler)
	// how use POST
	// curl -d '{ "task": "Go dev" }' -H "Content-Type: application/json" -X POST http://localhost:8080/users

	mux.HandleFunc("GET /users", getHandlers)
	// how use GET
	// curl http://localhost:8080/users

	fmt.Println("Server is working...")
	fmt.Println("http://localhost:8080/users")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
