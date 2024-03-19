package router

import (
	"context"

	"github.com/gin-gonic/gin"

	"github/kunhou/gl_exercise/internal/entity"
)

type TaskRouter struct {
	s ITaskService
}

//go:generate mockgen -source ./task_controller.go -destination=../mocks/task_controller.go -package=mocks
type ITaskService interface {
	List(ctx context.Context) (tasks []entity.Task, err error)
	Create(ctx context.Context, task entity.Task) (result entity.Task, err error)
	Update(ctx context.Context, id int, task entity.Task) (result entity.Task, err error)
	Delete(ctx context.Context, id int) error
}

func NewTaskRouter(s ITaskService) *TaskRouter {
	return &TaskRouter{
		s: s,
	}
}

func (t *TaskRouter) RegisterRouter(r *gin.Engine) {
	r.GET("/tasks", t.List)
}

// List lists tasks
// @Summary  List tasks
// @Tags     Task
// @Accept   json
// @Produce  json
// @Success  200  {object}  schema.Response{result=[]entity.Task}  "ok"
// @Router   /tasks [get]
func (t *TaskRouter) List(ctx *gin.Context) {
}

// Create creates a task
// @Summary  Create task
// @Tags     Task
// @Accept   json
// @Produce  json
// @Param    task  body      schema.TaskCreateRequest             true  "task"
// @Success  200   {object}  schema.Response{result=entity.Task}  "ok"
// @Failure  400   {object}  schema.Response{result=string}       "bad request"
// @Router   /tasks [post]
func (t *TaskRouter) Create(ctx *gin.Context) {
}

// Update updates a task
// @Summary  Update task
// @Tags     Task
// @Accept   json
// @Produce  json
// @Param    id   path      int                             true  "Task ID"
// @Param    task  body      schema.TaskUpdateRequest             true  "Update task"
// @Success  200   {object}  schema.Response{result=entity.Task}  "ok"
// @Failure  400   {object}  schema.Response{result=string}       "bad request"
// @Failure  404  {object}  schema.Response{result=string}  "not found"
// @Router   /tasks/{id} [put]
func (t *TaskRouter) Update(ctx *gin.Context) {
}

// Delete deletes a task
// @Summary  Delete task
// @Tags     Task
// @Accept   json
// @Produce  json
// @Param    id    path      int                                  true  "Task ID"
// @Success  200  {object}  schema.Response{result=string}  "ok"
// @Failure  404   {object}  schema.Response{result=string}       "not found"
// @Router   /tasks/{id} [delete]
func (t *TaskRouter) Delete(ctx *gin.Context) {
}
