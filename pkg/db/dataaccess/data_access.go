package dataaccess

import (
	"example.com/pkg/task"
)

type DataAccess interface {
	Get(id string) (task.Task, error)
	Put(id string, t task.Task) error
	Post(t task.Task) (string, error)
	Delete(id string) error
}
