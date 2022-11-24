package usecases

import (
	"errors"
	"reflect"
	"testing"
	task "todo/internal/task/model"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

type DeleteTaskRepositoryStub struct{}

func (t *DeleteTaskRepositoryStub) DeleteById(taskId int) error {
	return nil
}

func makeTask() *task.Task {
	task, _ := task.NewTask("new task", task.High)
	task.Id = 1
	return task
}
func makeDeleteTaskUseCase() *DeleteTaskUseCase {
	return NewDeleteTaskUseCase(&DeleteTaskRepositoryStub{})
}
func TestDeleteTaskUseCase_ReturnsErrorOnFail(t *testing.T) {
	sut := makeDeleteTaskUseCase()
	validAddTask := makeTask()
	monkey.PatchInstanceMethod(
		reflect.TypeOf(sut.DeleteTaskRepository),
		"DeleteById",
		func(t *DeleteTaskRepositoryStub, _ int) error {
			return errors.New("error on delete task")
		})
	err := sut.Execute(validAddTask.Id)
	assert.Error(t, err, "error on delete task")
	monkey.UnpatchInstanceMethod(reflect.TypeOf(sut.DeleteTaskRepository),"DeleteById")
}
func TestDeleteTaskUseCase(t *testing.T) {
	sut := makeDeleteTaskUseCase()
	validAddTask := makeTask()
	err := sut.Execute(validAddTask.Id)
	assert.Nil(t, err)
}
