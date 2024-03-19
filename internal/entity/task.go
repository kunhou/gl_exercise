package entity

type Task struct {
	Id     int        `json:"id"`
	Name   string     `json:"name"`
	Status TaskStatus `json:"status"`
}

type TaskStatus int

const (
	TaskStatusIncomplete TaskStatus = iota
	TaskStatusComplete
)
