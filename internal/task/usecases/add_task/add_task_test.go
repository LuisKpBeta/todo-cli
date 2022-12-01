package usecases

import (
	"errors"
	"reflect"
	"testing"
	task "todo/internal/task/model"
	dto "todo/internal/task/usecases/task_dto"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

type TaskRepositoryStub struct{}

func (t *TaskRepositoryStub) AddTask(task *task.Task) error {
	task.Id = 1
	return nil
}

func makeSut() *AddTaskUseCase {
	return NewAddTaskUseCase(&TaskRepositoryStub{})
}
func makeValidTask() dto.AddTaskDTO {
	addTaskDTO := dto.AddTaskDTO{
		Description: "new task",
		Priority:    "normal",
	}
	return addTaskDTO
}
func TestAddTaskCreateAndReturnNewTask(t *testing.T) {
	sut := makeSut()
	validAddTask := makeValidTask()
	validAddTask.Priority = "high"
	newTask, err := sut.Execute(validAddTask)
	assert.Nil(t, err)
	assert.Equal(t, newTask.Id, 1)
	assert.Equal(t, newTask.Description, validAddTask.Description)
	assert.Equal(t, newTask.Priority(),  task.High)
}
func TestAddTaskCreateTaskAsNormalPriorityWhenEmpty(t *testing.T) {
	sut := makeSut()
	validAddTask := makeValidTask()
	validAddTask.Priority = ""
	newTask, err := sut.Execute(validAddTask)
	assert.Nil(t, err)
	assert.Equal(t, newTask.Id, 1)
	assert.Equal(t, newTask.Description, validAddTask.Description)
	assert.Equal(t, newTask.Priority(), task.Normal)
}
func TestReturnErrorOnRepositoryReturnsError(t *testing.T) {
	sut := makeSut()
	monkey.PatchInstanceMethod(
		reflect.TypeOf(sut.TaskRepository),
		"AddTask",
		func(t *TaskRepositoryStub, _ *task.Task) error {
			return errors.New("error on insert new task")
		})
	validAddTask := makeValidTask()
	newTask, err := sut.Execute(validAddTask)
	assert.Error(t, err, "error on insert new task")
	assert.Nil(t, newTask)
}
