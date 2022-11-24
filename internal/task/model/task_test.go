package model

import (
	"os"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2022, 11, 19, 12, 59, 59, 0, time.Local)
	})
	m.Run()
	os.Exit(0)
}

func TestShouldCreateNewTask(t *testing.T) {
	newTask, _ := NewTask("nova task", Low)
	assert.Equal(t, newTask.Description, "nova task")
	assert.Equal(t, newTask.Priority, Low)
	assert.Equal(t, newTask.Status, false)
	assert.Equal(t, newTask.Created, time.Now())
}
func TestShouldReturnErrorOnInvalidDescription(t *testing.T) {
	_, err := NewTask("", Low)
	assert.Error(t, err, "descption must have at least 3 characters")
}
func TestShouldReturnErrorOnInvalidPriority(t *testing.T) {
	_, err := NewTask("new task", "more or less")
	assert.Error(t, err, "status must be low, high or normal")
}