package service

type StatementService interface {
	CreateTask(task string) (RequestBody, error)
	GetAllTasks() ([]RequestBody, error)
	GetTaskById(id uint) (RequestBody, error)
	UpdateTask(id uint, task string, isDone bool) (RequestBody, error)
	DeleteTask(id uint) error
}

type TaskService struct {
	repo StatementRepository
}

func NewTaskService(r StatementRepository) StatementService {
	return &TaskService{repo: r}
}

func (s *TaskService) CreateTask(task string) (RequestBody, error) {
	var act RequestBody

	act.Task = task

	if err := s.repo.CreateTask(&act); err != nil {
		return RequestBody{}, err
	}
	return act, nil
}

func (s *TaskService) GetAllTasks() ([]RequestBody, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) GetTaskById(id uint) (RequestBody, error) {
	return s.repo.GetTaskById(id)
}

func (s *TaskService) UpdateTask(id uint, task string, isDone bool) (RequestBody, error) {
	act, err := s.repo.GetTaskById(id)
	if err != nil {
		return RequestBody{}, err
	}

	act.Task = task
	act.IsDone = isDone

	if err := s.repo.UpdateTask(act); err != nil {
		return RequestBody{}, err
	}
	return act, nil
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTask(id)
}
