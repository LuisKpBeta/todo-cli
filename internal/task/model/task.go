package model

import (
	"errors"
	"time"
)

type TaskPriority int

const (
	Low    TaskPriority = 2
	Normal TaskPriority = 1
	High   TaskPriority = 0
)

type Task struct {
	Id          int
	Description string
	Status      bool
	Priority    TaskPriority
	Created     time.Time
}

func NewTask(description string, priority TaskPriority) (*Task, error) {
	task := &Task{
		Description: description,
		Status:      false,
		Priority:    priority,
		Created:     time.Now(),
	}
	isValid := task.IsValid()
	if isValid != nil {
		return nil, isValid
	}
	return task, nil

}
func (t *Task) IsValid() error {
	if len(t.Description) < 3 {
		return errors.New("descption must have at least 3 characters")
	}
	isValisPriority := t.IsValidPriority()
	if isValisPriority != nil {
		return isValisPriority
	}
	return nil
}
func (t *Task) IsValidPriority() error {
	switch t.Priority {
	case Low, Normal, High:
		return nil
	}
	return errors.New("status must be low, high or normal")
}
func (t *Task) CompleteTask() {
	t.Status = true
}
