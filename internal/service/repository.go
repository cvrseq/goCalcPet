package service

import (
	"gorm.io/gorm"
)

type StatementRepository interface {
	CreateTask(act *RequestBody) error
	GetAllTasks() ([]RequestBody, error)
	GetTaskById(id uint) (RequestBody, error)
	UpdateTask(task RequestBody) error
	DeleteTask(id uint) error
}

type StateRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) StatementRepository {
	return &StateRepository{db: db}
}

func (r *StateRepository) CreateTask(act *RequestBody) error {
	return r.db.Create(act).Error
}

func (r *StateRepository) GetAllTasks() ([]RequestBody, error) {
	var task []RequestBody

	return task, r.db.Find(&task).Error
}

func (r *StateRepository) GetTaskById(id uint) (RequestBody, error) {
	var task RequestBody
	err := r.db.First(&task, id).Error

	return task, err
}

func (r *StateRepository) UpdateTask(task RequestBody) error {
	return r.db.Save(&task).Error
}

func (r *StateRepository) DeleteTask(id uint) error {
	return r.db.Delete(&RequestBody{}, id).Error
}
