package task

import (
	"context"

	"github/kunhou/gl_exercise/internal/entity"
)

//go:generate mockgen -source ./service.go -destination=../mocks/task.go -package=mocks
// ITaskRepository task repository
type ITaskRepository interface {
	List(ctx context.Context) ([]entity.Task, error)
	Create(ctx context.Context, task entity.Task) (entity.Task, error)
	Update(ctx context.Context, id int, task entity.Task) (entity.Task, error)
	Delete(ctx context.Context, id int) error
}

// Service task service
type Service struct {
	taskRepo ITaskRepository
}

// New new task service
func New(taskRepo ITaskRepository) *Service {
	return &Service{
		taskRepo: taskRepo,
	}
}
