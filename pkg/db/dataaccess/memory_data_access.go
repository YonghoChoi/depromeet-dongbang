package dataaccess

import (
	"errors"
	"example.com/pkg/task"
	"fmt"
)

var ErrTaskNotExist = errors.New("task does not exist")

type MemoryDataAccess struct {
	tasks  map[ID]task.Task
	nextID int64
}

func (m *MemoryDataAccess) Get(id ID) (task.Task, error) {
	t, exists := m.tasks[id]
	if !exists {
		return task.Task{}, ErrTaskNotExist
	}

	return t, nil
}

func (m *MemoryDataAccess) Put(id ID, t task.Task) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
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

func (m *MemoryDataAccess) Delete(id ID) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
	}
	delete(m.tasks, id)
	return nil
}

func NewMemoryDataAccess() DataAccess {
	return &MemoryDataAccess{
		tasks:  map[ID]task.Task{},
		nextID: int64(1),
	}
}
