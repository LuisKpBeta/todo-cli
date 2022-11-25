package usecases

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
	task "todo/internal/task/model"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

type ListTaskRepositoryStub struct {
	CalledArgs struct {
		listAll         bool
		orderByPriority bool
	}
}

func (t *ListTaskRepositoryStub) ListTasks(listAll bool, orderByPriority bool) ([]task.Task, error) {
	t.CalledArgs = struct {
		listAll         bool
		orderByPriority bool
	}{
		listAll:         listAll,
		orderByPriority: orderByPriority,
	}
	var taskList []task.Task
	validAddTask := task.Task{
		Id:          1,
		Description: "new task",
		Status:      false,
		Priority:    task.High,
		Created:     time.Now(),
	}
	taskList = append(taskList, validAddTask)
	return taskList, nil
}

func makeListTaskUseCase() (*ListTaskUseCase, *ListTaskRepositoryStub) {
	repoStub := &ListTaskRepositoryStub{}
	return NewListTaskUseCase(repoStub), repoStub
}
func TestListTaskUseCase_ReturnsErrorOnFail(t *testing.T) {
	sut, _ := makeListTaskUseCase()
	monkey.PatchInstanceMethod(
		reflect.TypeOf(sut.TaskRepository),
		"ListTasks",
		func(t *ListTaskRepositoryStub, _ bool, _ bool) ([]task.Task, error) {
			return nil, errors.New("error on load tasks")
		})
	_, err := sut.Execute(ListTaskArgs{})
	fmt.Println(err)
	assert.Error(t, err, "error on load task")
	monkey.UnpatchInstanceMethod(reflect.TypeOf(sut.TaskRepository), "ListTasks")
}
func TestListTaskUseCase(t *testing.T) {
	sut, _ := makeListTaskUseCase()
	tasks, err := sut.Execute(ListTaskArgs{})
	assert.Nil(t, err)
	assert.Equal(t, len(tasks), 1)
	assert.Equal(t, tasks[0].Id, 1)
	assert.Equal(t, tasks[0].Description, "new task")
}
func TestListTaskUseCaseCallsRepositoryWithCorrectArgs(t *testing.T) {
	sut, repoSpy := makeListTaskUseCase()
	_, err := sut.Execute(ListTaskArgs{ListAll: false, OrderByPriority: true})
	assert.Nil(t, err)
	assert.True(t, repoSpy.CalledArgs.orderByPriority)
	assert.False(t, repoSpy.CalledArgs.listAll)

}
