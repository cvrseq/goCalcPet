package main

import (
	"CRUD/internal/db"
	"CRUD/internal/handlers"
	"CRUD/internal/service"
	"CRUD/internal/web/tasks"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Could not connect database", err)
	}

	//router := http.NewServeMux()

	taskRepo := service.NewTaskRepository(database)
	taskService := service.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//router.HandleFunc("POST /tasks", taskHandlers.CreateHandler)
	//router.HandleFunc("GET /tasks", taskHandlers.GetHandlers)
	//router.HandleFunc("PATCH /tasks/{id}", taskHandlers.UpdateHandler)
	//router.HandleFunc("DELETE /tasks/{id}", taskHandlers.DeleteHandler)

	fmt.Println("Server is working...")
	fmt.Println("http://localhost:8080/tasks")
	//log.Fatal(http.ListenAndServe(":8080", router))
	strictHandler := tasks.NewStrictHandler(taskHandlers, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}

}
