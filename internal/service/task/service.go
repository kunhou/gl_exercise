package task

type ITaskRepository interface {
}

type Service struct {
	taskRepo ITaskRepository
}

func New(taskRepo ITaskRepository) *Service {
	return &Service{
		taskRepo: taskRepo,
	}
}
