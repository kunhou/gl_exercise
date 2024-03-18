package task

import (
	"context"

	"github/kunhou/gl_exercise/internal/entity"
	"github/kunhou/gl_exercise/internal/service/task"
)

type TaskRepo struct {
}

// new setting repository
func NewTaskRepo() task.ITaskRepository {
	return &TaskRepo{}
}

func (t *TaskRepo) List(ctx context.Context) (tasks []entity.Task, err error) {
	return
}

func (t *TaskRepo) Create(ctx context.Context, task entity.Task) (result entity.Task, err error) {
	return
}

func (t *TaskRepo) Update(ctx context.Context, id int, task entity.Task) (result entity.Task, err error) {
	return
}

func (t *TaskRepo) Delete(ctx context.Context, id int) error {
	return nil
}
