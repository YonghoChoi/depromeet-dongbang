package dataaccess

import (
	"errors"
	"example.com/model/code"
	"example.com/pkg/task"
	"fmt"
)

type MemoryDataAccess struct {
	tasks  map[string]task.Task
	nextID int64
}

func (m *MemoryDataAccess) Get(id string) (task.Task, error) {
	t, exists := m.tasks[id]
	if !exists {
		return task.Task{}, code.ErrTaskNotExist
	}

	return t, nil
}

func (m *MemoryDataAccess) Put(id string, t task.Task) error {
	if _, exists := m.tasks[id]; !exists {
		return code.ErrTaskNotExist
	}

	m.tasks[id] = t
	return nil
}

func (m *MemoryDataAccess) Post(t task.Task) (ID, error) {
	id := ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.tasks[id] = t
	return id, nil
}

func (m *MemoryDataAccess) Delete(id string) error {
	if _, exists := m.tasks[id]; !exists {
		return code.ErrTaskNotExist
	}
	delete(m.tasks, id)
	return nil
}

func NewMemoryDataAccess() DataAccess {
	return &MemoryDataAccess{
		tasks:  map[string]task.Task{},
		nextID: int64(1),
	}
}
