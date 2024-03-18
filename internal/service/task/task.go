package task

import (
	"context"

	"github/kunhou/gl_exercise/internal/entity"
)

func (s *Service) List(ctx context.Context) (tasks []entity.Task, err error) {
	tasks, err = s.taskRepo.List(ctx)
	return
}

func (s *Service) Create(ctx context.Context, task entity.Task) (result entity.Task, err error) {
	result, err = s.taskRepo.Create(ctx, task)
	return
}

func (s *Service) Update(ctx context.Context, id int, task entity.Task) (result entity.Task, err error) {
	return
}
