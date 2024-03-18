package task

import (
	"context"

	"github/kunhou/gl_exercise/internal/entity"
)

// List list tasks
func (s *Service) List(ctx context.Context) (tasks []entity.Task, err error) {
	tasks, err = s.taskRepo.List(ctx)
	return
}

// Create create task
func (s *Service) Create(ctx context.Context, task entity.Task) (result entity.Task, err error) {
	result, err = s.taskRepo.Create(ctx, task)
	return
}

// Update update task
func (s *Service) Update(ctx context.Context, id int, task entity.Task) (result entity.Task, err error) {
	result, err = s.taskRepo.Update(ctx, id, task)
	return
}

// Delete delete task
func (s *Service) Delete(ctx context.Context, id int) error {
	return s.taskRepo.Delete(ctx, id)
}
