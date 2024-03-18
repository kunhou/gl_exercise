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

	// auto increment id
	lastID int
}

// NewTaskRepo new task repository
func NewTaskRepo() task.ITaskRepository {
	return &TaskRepo{}
}

// List list tasks
func (t *TaskRepo) List(ctx context.Context) (tasks []entity.Task, err error) {
	tasks = make([]entity.Task, 0)

	t.rwmu.RLock()
	defer t.rwmu.RUnlock()

	for _, task := range t.taskStorage {
		tasks = append(tasks, task)
	}

	return
}

// Create create task
func (t *TaskRepo) Create(ctx context.Context, task entity.Task) (result entity.Task, err error) {
	task.Id = t.autoIncrementID()

	t.rwmu.Lock()
	defer t.rwmu.Unlock()

	t.taskStorage[task.Id] = task

	result = task
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

func (t *TaskRepo) autoIncrementID() int {
	t.rwmu.Lock()
	defer t.rwmu.Unlock()

	t.lastID++

	return t.lastID
}
