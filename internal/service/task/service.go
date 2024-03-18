package task

import (
	"context"

	"github/kunhou/gl_exercise/internal/entity"
)

//go:generate mockgen -source ./service.go -destination=../mocks/task.go -package=mocks
type ITaskRepository interface {
	List(ctx context.Context) ([]entity.Task, error)
}

type Service struct {
	taskRepo ITaskRepository
}

func New(taskRepo ITaskRepository) *Service {
	return &Service{
		taskRepo: taskRepo,
	}
}
