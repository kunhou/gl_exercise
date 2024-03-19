package router

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github/kunhou/gl_exercise/internal/common/reason"
	"github/kunhou/gl_exercise/internal/deliver/http/schema"
	"github/kunhou/gl_exercise/internal/entity"
	merr "github/kunhou/gl_exercise/internal/pkg/errors"
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
	r.POST("/task", t.Create)
	r.PUT("/task/:id", t.Update)
	r.DELETE("/task/:id", t.Delete)
}

// List lists tasks
// @Summary  List tasks
// @Tags     Task
// @Accept   json
// @Produce  json
// @Success  200  {object}  schema.Response{result=[]entity.Task}  "ok"
// @Router   /tasks [get]
func (t *TaskRouter) List(ctx *gin.Context) {
	tasks, err := t.s.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schema.Response{
			Result: reason.UnknownError,
		})
		return
	}

	ctx.JSON(http.StatusOK, schema.Response{
		Result: tasks,
	})
}

// Create creates a task
// @Summary  Create task
// @Tags     Task
// @Accept   json
// @Produce  json
// @Param    task  body      schema.TaskCreateRequest             true  "task"
// @Success  200   {object}  schema.Response{result=entity.Task}  "ok"
// @Failure  400   {object}  schema.Response{result=string}       "bad request"
// @Router   /task [post]
func (t *TaskRouter) Create(ctx *gin.Context) {
	req := &schema.TaskCreateRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, schema.Response{
			Result: reason.RequestFormatError,
		})
		return
	}

	task := entity.Task{
		Name: req.Name,
	}

	result, err := t.s.Create(ctx, task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schema.Response{
			Result: reason.UnknownError,
		})
		return
	}

	ctx.JSON(http.StatusCreated, schema.Response{
		Result: result,
	})
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
// @Router   /task/{id} [put]
func (t *TaskRouter) Update(ctx *gin.Context) {
	req := &schema.TaskUpdateRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, schema.Response{
			Result: reason.RequestFormatError,
		})
		return
	}

	param := &schema.TaskParam{}
	if err := ctx.BindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, schema.Response{
			Result: reason.RequestFormatError,
		})
		return
	}

	var status entity.TaskStatus
	switch *req.Status {
	case 0:
		status = entity.TaskStatusIncomplete
	case 1:
		status = entity.TaskStatusComplete
	}

	task := entity.Task{
		Name:   req.Name,
		Status: status,
	}

	result, err := t.s.Update(ctx, param.Id, task)
	if err != nil {
		var myErr *merr.Error
		if errors.As(err, &myErr) {
			ctx.JSON(myErr.Code, schema.Response{
				Result: myErr.Reason,
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, schema.Response{
			Result: reason.UnknownError,
		})
		return
	}

	ctx.JSON(http.StatusOK, schema.Response{
		Result: result,
	})
}

// Delete deletes a task
// @Summary  Delete task
// @Tags     Task
// @Accept   json
// @Produce  json
// @Param    id    path      int                                  true  "Task ID"
// @Success  200  {object}  schema.Response{result=string}  "ok"
// @Failure  404   {object}  schema.Response{result=string}       "not found"
// @Router   /task/{id} [delete]
func (t *TaskRouter) Delete(ctx *gin.Context) {
	param := &schema.TaskParam{}
	if err := ctx.BindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, schema.Response{
			Result: reason.RequestFormatError,
		})
		return
	}

	err := t.s.Delete(ctx, param.Id)
	if err != nil {
		var myErr *merr.Error
		if errors.As(err, &myErr) {
			ctx.JSON(myErr.Code, schema.Response{
				Result: myErr.Reason,
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, schema.Response{
			Result: reason.UnknownError,
		})
		return
	}

	ctx.JSON(http.StatusOK, schema.Response{
		Result: reason.Success,
	})
}
