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
	}

	//if r.Method != r.PostFormValue(task) {
	//	fmt.Printf("%v, \n", http.StatusBadRequest)
	//}

	fmt.Println("successful decode")

	task = rq.Task

	fmt.Println("json data write in the task variable")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/endpoint", createHandler)

	fmt.Println("Server is working...")
	fmt.Println("http://localhost:8080/endpoint")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
