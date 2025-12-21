package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type requestBody struct {
	Task string `json:"task"`
	ID   string `json:"id"`
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

func updateHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	fmt.Fprintf(w, "%s", id)

	fmt.Println(id)

	var rq requestBody

	if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
		fmt.Printf("%v", "error not decode from Body to reqBody struct")
		return
	}

	fmt.Println("successful decode")

	task = rq.Task

	fmt.Println("json data update in the task variable")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var rq requestBody

	if id == rq.ID {
		task = " "
	}

	if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
		fmt.Printf("%v", "error not decode from Body to reqBody struct")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("POST /users", createHandler)
	// how use POST
	// curl -d '{ "task": "Go dev" }' -H "Content-Type: application/json" -X POST http://localhost:8080/users

	router.HandleFunc("GET /users", getHandlers)
	// how use GET
	// curl http://localhost:8080/users

	router.HandleFunc("PATCH /users/{id}", updateHandler)
	//how use PATCH
	// curl -X PATCH http://localhost:8080/users/{id} -H "Content-Type: application/json" -d '{"task": "rust developer"}'

	router.HandleFunc("DELETE /users/{id}", deleteHandler)
	//how use DELETE
	// curl -X DELETE http://localhost:8080/users/{id}

	fmt.Println("Server is working...")
	fmt.Println("http://localhost:8080/users")
	log.Fatal(http.ListenAndServe(":8080", router))

}
