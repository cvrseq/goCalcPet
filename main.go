package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var db *gorm.DB

func initDB() {
	dataSrcName := "host=localhost user=postgres password=yourpassword dbname=postgres port=5000 sslmode=disable"

	var err error

	db, err = gorm.Open(postgres.Open(dataSrcName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := db.AutoMigrate(&requestBody{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}

type requestBody struct {
	Task   string `gorm:"column:task" json:"task"`
	ID     uint   `gorm:"primaryKey" json:"id"`
	IsDone bool   `gorm:"column:is_done" json:"is_done"`
}

func createHandler(w http.ResponseWriter, r *http.Request) {

	var task requestBody

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Could not decode in the create", http.StatusBadRequest)
		return
	}

	fmt.Println("successful decode")

	task.ID = 0

	fmt.Println("Trying create variable and write db")
	if err := db.Create(&task).Error; err != nil {
		http.Error(w, "Could not create in db", http.StatusInternalServerError)
		return
	}

	fmt.Println("json data write in the task variable")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Could not decode in the create", http.StatusNotFound)
		return
	}

}

func getHandlers(w http.ResponseWriter, r *http.Request) {
	var task []requestBody

	if err := db.Find(&task).Error; err != nil {
		http.Error(w, "Could not find on db", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Could not find on db", http.StatusNotFound)
		return
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	var task requestBody
	id := r.PathValue("id")

	var rq struct {
		Task   *string `json:"task"`
		IsDone *bool   `json:"is_done"`
	}

	if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
		http.Error(w, "Could not decode in the update", http.StatusBadRequest)
		return
	}

	fmt.Println("successful decode")

	if err := db.First(&task, "id = ?", id).Error; err != nil {
		http.Error(w, "Could not first on db", http.StatusInternalServerError)
		return
	}

	if rq.Task != nil {
		task.Task = *rq.Task
	}

	if rq.IsDone != nil {
		task.IsDone = *rq.IsDone
	}

	if err := db.Save(&task).Error; err != nil {
		http.Error(w, "Could not save on db", http.StatusInternalServerError)
		return
	}

	fmt.Println("json data update in the task variable")

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Could not encode in the update", http.StatusNotFound)
		return
	}

}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	fmt.Println("successful decode in the delete")

	if err := db.Delete(&requestBody{}, id).Error; err != nil {
		http.Error(w, "Could not delete on db", http.StatusInternalServerError)
		return

	}

	w.WriteHeader(http.StatusNoContent)

}

func main() {
	initDB()
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
