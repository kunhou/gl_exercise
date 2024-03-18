package router

import "github/kunhou/gl_exercise/internal/service/task"

type TaskRouter struct {
	s *task.Service
}

func NewTaskRouter(s *task.Service) *TaskRouter {
	return &TaskRouter{
		s: s,
	}
}
