package usecases

import (
	"errors"
	"reflect"
	"testing"
	task "todo/internal/task/model"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

type CompleteTaskRepositoryStub struct{}

func (t *CompleteTaskRepositoryStub) CompleteTask(task *task.Task) error {
	task.Status = true
	return nil
}
func (t *CompleteTaskRepositoryStub) FindById(taskId int) (*task.Task, error) {
	task := makeTask()
	return task, nil
}
func makeTask() *task.Task {
	task, _ := task.NewTask("new task", task.High)
	task.Id = 1
	return task
}
func makeCompleteTaskUseCase() *CompleteTaskUseCase {
	return NewCompleteTaskUseCase(&CompleteTaskRepositoryStub{}, &CompleteTaskRepositoryStub{})
}
func TestCompleteTaskUseCase(t *testing.T) {
	sut := makeCompleteTaskUseCase()
	validAddTask := makeTask()
	task, err := sut.Execute(validAddTask.Id)
	assert.Nil(t, err)
	assert.Equal(t, task.Id, 1)
	assert.True(t, task.Status)
}
func TestCompleteTaskUseCase_WhenFail(t *testing.T) {
	sut := makeCompleteTaskUseCase()
	validAddTask := makeTask()
	monkey.PatchInstanceMethod(
		reflect.TypeOf(sut.CompleteTaskRepository),
		"CompleteTask",
		func(t *CompleteTaskRepositoryStub, _ *task.Task) error {
			return errors.New("error on update task")
		})
	task, err := sut.Execute(validAddTask.Id)
	assert.Error(t, err, "error on update task")
	assert.Nil(t, task)
}
func TestCompleteTaskUseCase_WhenNotFindTask(t *testing.T) {
	sut := makeCompleteTaskUseCase()
	validAddTask := makeTask()
	monkey.PatchInstanceMethod(
		reflect.TypeOf(sut.FindTaskRepository),
		"FindById",
		func(t *CompleteTaskRepositoryStub, _ int) (*task.Task, error) {
			return nil, nil
		})
	task, err := sut.Execute(validAddTask.Id)
	assert.Error(t, err, "error on update task")
	assert.Nil(t, task)
}
