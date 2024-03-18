package task

import (
	"context"
	"sync"

	"github/kunhou/gl_exercise/internal/entity"
	"github/kunhou/gl_exercise/internal/service/task"
)

// TaskRepo task repository
type TaskRepo struct {
	taskStorage map[int]entity.Task
	rwmu        sync.RWMutex
}

// NewTaskRepo new task repository
func NewTaskRepo() task.ITaskRepository {
	return &TaskRepo{}
}

// List list tasks
func (t *TaskRepo) List(ctx context.Context) (tasks []entity.Task, err error) {
	return
}

// Create create task
func (t *TaskRepo) Create(ctx context.Context, task entity.Task) (result entity.Task, err error) {
	return
}

// Update update task
func (t *TaskRepo) Update(ctx context.Context, id int, task entity.Task) (result entity.Task, err error) {
	return
}

// Delete delete task
func (t *TaskRepo) Delete(ctx context.Context, id int) error {
	return nil
}
