package main

import (
	"net/http"
)

type requestBody struct {
	Task string `json:"task"`
}

var task string

func createHandler(w http.ResponseWriter, r *http.Request) error {
	// pass
}

func main() {

}
