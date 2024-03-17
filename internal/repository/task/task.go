package task

import "github/kunhou/gl_exercise/internal/service/task"

type TaskRepo struct {
}

// new setting repository
func NewTaskRepo() task.ITaskRepository {
	return &TaskRepo{}
}
