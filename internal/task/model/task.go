package model

import (
	"errors"
	"fmt"
	"strings"
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
func (t *Task) Age() string {
	format := func(txt string, value int) string {
		if value > 1 {
			return fmt.Sprint(txt, "s")
		}
		return txt
	}
	now := time.Now()
	created := t.Created
	days := int(now.Sub(created).Abs().Hours() / 24)
	var age_string string
	if days >= 1 {
		d_text := format("day", days)
		created = created.AddDate(0, 0, days)
		age_string = fmt.Sprint(days, " ", d_text)
	}
	age := now.Sub(created)

	if age.Minutes() != 0 {
		hours := int(age.Minutes() / 60)
		minutes := int(age.Minutes()) % 60
		if hours > 0 {
			h_text := format("hour", hours)
			age_string = fmt.Sprint(age_string, " ", hours, " ", h_text)
		}
		m_text := format("minute", minutes)
		age_string = fmt.Sprint(age_string, " ", minutes, " ", m_text)
	}
	return strings.TrimSpace(age_string)
}
func (t *Task) PriorityToString() string {
	switch t.Priority {
	case Low:
		return "low"
	case Normal:
		return "normal"
	case High:
		return "high"
	default:
		return ""
	}
	
}
