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
		log.Fatal(err)
		return
	}

	fmt.Println("successful decode")

	task = rq.Task

	fmt.Println("json data write in the task variable")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

func getHandlers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(task); err != nil {
		log.Fatal(err)
		return
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if n, err := fmt.Fprintf(w, "%s", id); err != nil {
		fmt.Println("not change var", n)
	}

	fmt.Println(id)

	var rq requestBody

	if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("successful decode")

	task = rq.Task

	fmt.Println("json data update in the task variable")

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(task); err != nil {
		log.Fatal(err)
		return
	}

}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var rq requestBody

	fmt.Println("successful decode in the delete")

	if id == rq.ID {
		task = ""
	}

	w.WriteHeader(http.StatusNoContent)

}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("POST /tasks", createHandler)
	// how use POST
	// curl -d '{ "task": "Go dev" }' -H "Content-Type: application/json" -X POST http://localhost:8080/tasks

	router.HandleFunc("GET /tasks", getHandlers)
	// how use GET
	// curl http://localhost:8080/tasks

	router.HandleFunc("PATCH /tasks/{id}", updateHandler)
	//how use PATCH
	// curl -X PATCH http://localhost:8080/tasks/{id} -H "Content-Type: application/json" -d '{"task": "rust developer"}'

	router.HandleFunc("DELETE /tasks/{id}", deleteHandler)
	//how use DELETE
	// curl -X DELETE http://localhost:8080/tasks/{id}

	fmt.Println("Server is working...")
	fmt.Println("http://localhost:8080/tasks")
	log.Fatal(http.ListenAndServe(":8080", router))

}
