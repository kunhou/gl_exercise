package schema

type TaskCreateRequest struct {
	Name string `json:"name" example:"task 1" binding:"required"`
}

type TaskParam struct {
	Id int `swaggerignore:"true" uri:"id" form:"id" binding:"required"`
}

type TaskUpdateRequest struct {
	Name   string `json:"name" example:"task 1" binding:"required"`
	Status *int   `json:"status" example:"1" binding:"required,oneof=0 1"`
}
