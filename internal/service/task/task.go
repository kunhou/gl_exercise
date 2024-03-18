package task

import (
	"context"

	"github/kunhou/gl_exercise/internal/entity"
)

func (s *Service) List(ctx context.Context) (tasks []entity.Task, err error) {
	tasks, err = s.taskRepo.List(ctx)
	return
}
