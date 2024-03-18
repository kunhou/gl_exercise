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
