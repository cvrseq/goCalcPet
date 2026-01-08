package handlers

import (
	"CRUD/internal/service"
	"CRUD/internal/web/tasks"
	"context"
	"fmt"
)

type StatementHandler struct {
	service service.StatementService
}

func NewTaskHandler(s service.StatementService) *StatementHandler {
	return &StatementHandler{service: s}
}

func (h *StatementHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {

		idVal := int(tsk.ID)

		task := tasks.Task{
			Id:     &idVal,
			IsDone: &tsk.IsDone,
			Task:   &tsk.Task,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *StatementHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	if request.Body == nil {
		return nil, fmt.Errorf("request body is required")
	}
	taskRequest := request.Body

	taskToCreate := service.RequestBody{}

	if taskRequest.Task != nil {
		taskToCreate.Task = *taskRequest.Task
	} else {
		return nil, fmt.Errorf("field 'task' is required")
	}

	if taskRequest.IsDone != nil {
		taskToCreate.IsDone = *taskRequest.IsDone
	}

	createdTask, err := h.service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	crTsk := int(createdTask.ID)

	response := tasks.PostTasks201JSONResponse{
		Id:     &crTsk,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}

	return response, nil
}

// /
// / COMPLETE THIS METHOD
// /
func (h *StatementHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {

	if request.Body == nil {
		return nil, fmt.Errorf("request body is required")
	}

	id := uint(request.Id)

	taskToUpdate := service.RequestBody{}

	taskRequest := request.Body

	if taskRequest.Task != nil {
		taskToUpdate.Task = *taskRequest.Task
	} else {
		return nil, fmt.Errorf("field 'task' is required")
	}

	var isBool bool
	if taskRequest.IsDone != nil {
		isBool = *taskRequest.IsDone
	}

	updatedTask, err := h.service.UpdateTask(id, taskToUpdate, isBool)
	if err != nil {
		return nil, err
	}

	crTsk := int(updatedTask.ID)

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &crTsk,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}
	return response, nil
}

// /
// / COMPLETE THIS METHOD
// /
func (h *StatementHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

//func (h *StatementHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
//
//	var task service.RequestBody
//
//	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
//		http.Error(w, "Could not decode in the create", http.StatusBadRequest)
//		return
//	}
//
//	act, err := h.service.CreateTask(task.Task)
//	if err != nil {
//		http.Error(w, "Could not decode in the create", http.StatusInternalServerError)
//		return
//	}
//
//	fmt.Println("successful decode")
//	fmt.Println("Trying create variable and write db...")
//	fmt.Println("json data write in the task variable")
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusCreated)
//	if err := json.NewEncoder(w).Encode(act); err != nil {
//		http.Error(w, "Could not decode in the create", http.StatusNotFound)
//		return
//	}
//
//}

//func (h *StatementHandler) GetHandlers(w http.ResponseWriter, r *http.Request) {
//	task, err := h.service.GetAllTasks()
//	if err != nil {
//		http.Error(w, "Could not find on db", http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//
//	if err := json.NewEncoder(w).Encode(task); err != nil {
//		http.Error(w, "Could not find on db", http.StatusNotFound)
//		return
//	}
//}

//func (h *StatementHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
//	strID := r.PathValue("id")
//
//	resID, err := strconv.Atoi(strID)
//	if err != nil {
//		http.Error(w, "invalid id format", http.StatusBadRequest)
//		return
//	}
//
//	id := uint(resID)
//
//	var rq struct {
//		Task   string `json:"task"`
//		IsDone bool   `json:"is_done"`
//	}
//
//	if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
//		http.Error(w, "Could not decode in the update", http.StatusBadRequest)
//		return
//	}
//
//	act, err := h.service.UpdateTask(id, rq.Task, rq.IsDone)
//	if err != nil {
//		http.Error(w, "Could not encode in the update", http.StatusInternalServerError)
//		return
//	}
//
//	fmt.Println("successful decode")
//	fmt.Println("json data update in the task variable")
//
//	w.Header().Set("Content-Type", "application/json")
//
//	if err := json.NewEncoder(w).Encode(act); err != nil {
//		http.Error(w, "Could not encode in the update", http.StatusNotFound)
//		return
//	}
//
//}
//
//func (h *StatementHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
//	strID := r.PathValue("id")
//
//	resID, err := strconv.Atoi(strID)
//	if err != nil {
//		http.Error(w, "invalid id format", http.StatusBadRequest)
//		return
//	}
//
//	id := uint(resID)
//
//	err = h.service.DeleteTask(id)
//	if err != nil {
//		http.Error(w, "Could not encode in the update", http.StatusInternalServerError)
//		return
//	}
//	fmt.Println("successful decode in the delete")
//
//	w.WriteHeader(http.StatusNoContent)
//
//}
