package main

import (
	"CRUD/internal/db"
	"CRUD/internal/handlers"
	"CRUD/internal/service"
	"fmt"
	"log"
	"net/http"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Could not connect database", err)
	}

	router := http.NewServeMux()

	taskRepo := service.NewTaskRepository(database)
	taskService := service.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)

	router.HandleFunc("POST /tasks", taskHandlers.CreateHandler)
	// how use POST
	// curl -d '{ "task": "Go dev" }' -H "Content-Type: application/json" -X POST http://localhost:8080/tasks

	router.HandleFunc("GET /tasks", taskHandlers.GetHandlers)
	// how use GET
	// curl http://localhost:8080/tasks

	router.HandleFunc("PATCH /tasks/{id}", taskHandlers.UpdateHandler)
	//how use PATCH
	// curl -X PATCH http://localhost:8080/tasks/{id} -H "Content-Type: application/json" -d '{"task": "rust developer"}'

	router.HandleFunc("DELETE /tasks/{id}", taskHandlers.DeleteHandler)
	//how use DELETE
	// curl -X DELETE http://localhost:8080/tasks/{id}

	fmt.Println("Server is working...")
	fmt.Println("http://localhost:8080/tasks")
	log.Fatal(http.ListenAndServe(":8080", router))

}
