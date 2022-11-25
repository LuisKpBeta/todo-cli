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
		return time.Date(2022, 11, 19, 12, 0, 0, 0, time.Local)
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
	_, err := NewTask("new task", 3)
	assert.Error(t, err, "status must be low, high or normal")
}

func TestShouldMarkTaskStatusAsTrue(t *testing.T) {
	newTask, _ := NewTask("nova task", Low)
	assert.Equal(t, newTask.Status, false)
	newTask.CompleteTask()
	assert.Equal(t, newTask.Status, true)
}

func TestReturnFormatedAgeOfTask(t *testing.T) {
	newTask, _ := NewTask("nova task", Low)
	newTask.Created = time.Date(2022, 11, 10, 12, 0, 0, 0, time.Local)
	age := newTask.Age()
	assert.Equal(t, age, "9 days")
	newTask.Created = time.Date(2022, 11, 18, 21, 30, 0, 0, time.Local)
	age = newTask.Age()
	assert.Equal(t, age, "14 hours 30 minutes")

}
func TestReturnPriorityInStringFormat(t *testing.T) {
	newTask, _ := NewTask("nova task", Low)
	priority := newTask.PriorityToString()
	assert.Equal(t, priority, "low")
	newTask.Priority = Normal
	priority = newTask.PriorityToString()
	assert.Equal(t, priority, "normal")
	newTask.Priority = High
	priority = newTask.PriorityToString()
	assert.Equal(t, priority, "high")
}
