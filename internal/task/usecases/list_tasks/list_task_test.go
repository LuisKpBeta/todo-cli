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
	}
}

func makeTaskList() []task.Task {
	var taskList []task.Task
	validAddTask := task.Task{
		Id:          1,
		Description: "new task",
		Status:      false,
		Created:     time.Date(2022, 11, 19, 12, 0, 0, 0, time.Local),
	}
	validAddTask.SetPriority(task.High)
	taskList = append(taskList, validAddTask)
	return taskList
}
func (t *ListTaskRepositoryStub) ListTasks(listAll bool) ([]task.Task, error) {
	t.CalledArgs = struct {
		listAll         bool
	}{
		listAll:         listAll,
	}
	taskList := makeTaskList()
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
		func(t *ListTaskRepositoryStub, _ bool) ([]task.Task, error) {
			return nil, errors.New("error on load tasks")
		})
	_, err := sut.Execute(false)
	fmt.Println(err)
	assert.Error(t, err, "error on load task")
	monkey.UnpatchInstanceMethod(reflect.TypeOf(sut.TaskRepository), "ListTasks")
}
func TestListTaskUseCase(t *testing.T) {
	sut, _ := makeListTaskUseCase()
	tasks, err := sut.Execute(false)
	assert.Nil(t, err)
	assert.Equal(t, len(tasks), 1)
	assert.Equal(t, tasks[0].Id, 1)
	assert.Equal(t, tasks[0].Description, "new task")
}
func TestListTaskUseCaseCallsRepositoryWithCorrectArgs(t *testing.T) {

	sut, repoSpy := makeListTaskUseCase()
	_, err := sut.Execute(true)
	assert.Nil(t, err)
	assert.True(t, repoSpy.CalledArgs.listAll)
}

func TestListTaskUseCaseReturnsListOfReadTaskDTO(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2022, 11, 19, 13, 25, 0, 0, time.Local)
	})
	sut, _ := makeListTaskUseCase()
	tasks, err := sut.Execute(true)
	assert.Nil(t, err)
	assert.Equal(t, len(tasks), 1)
	assert.Equal(t, tasks[0].Id, 1)
	assert.Equal(t, tasks[0].Age, "1 hour 25 minutes")
	assert.Equal(t, tasks[0].Priority, "high")
}

