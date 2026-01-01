package handlers

import (
	"CRUD/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type StatementHandler struct {
	service service.StatementService
}

func NewTaskHandler(s service.StatementService) *StatementHandler {
	return &StatementHandler{service: s}
}

func (h *StatementHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {

	var task service.RequestBody

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Could not decode in the create", http.StatusBadRequest)
		return
	}

	act, err := h.service.CreateTask(task.Task)
	if err != nil {
		http.Error(w, "Could not decode in the create", http.StatusInternalServerError)
		return
	}

	fmt.Println("successful decode")
	fmt.Println("Trying create variable and write db...")
	fmt.Println("json data write in the task variable")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(act); err != nil {
		http.Error(w, "Could not decode in the create", http.StatusNotFound)
		return
	}

}

func (h *StatementHandler) GetHandlers(w http.ResponseWriter, r *http.Request) {
	task, err := h.service.GetAllTasks()
	if err != nil {
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

func (h *StatementHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	strID := r.PathValue("id")

	resID, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return
	}

	id := uint(resID)

	var rq struct {
		Task   string `json:"task"`
		IsDone bool   `json:"is_done"`
	}

	if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
		http.Error(w, "Could not decode in the update", http.StatusBadRequest)
		return
	}

	act, err := h.service.UpdateTask(id, rq.Task, rq.IsDone)
	if err != nil {
		http.Error(w, "Could not encode in the update", http.StatusInternalServerError)
		return
	}

	fmt.Println("successful decode")
	fmt.Println("json data update in the task variable")

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(act); err != nil {
		http.Error(w, "Could not encode in the update", http.StatusNotFound)
		return
	}

}

func (h *StatementHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	strID := r.PathValue("id")

	resID, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return
	}

	id := uint(resID)

	err = h.service.DeleteTask(id)
	if err != nil {
		http.Error(w, "Could not encode in the update", http.StatusInternalServerError)
		return
	}
	fmt.Println("successful decode in the delete")

	w.WriteHeader(http.StatusNoContent)

}
