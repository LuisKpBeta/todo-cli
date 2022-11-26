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
func (t *ListTaskRepositoryStub) ListNextTasks() ([]task.Task, error) {
	taskList := makeTaskList()
	return taskList, nil
}

func makeListNextTaskUseCase() *ListNextTaskUseCase {
	repoStub := &ListTaskRepositoryStub{}
	return NewListNextTaskUseCase(repoStub)
}
func TestListNextTaskUseCase_ReturnsErrorOnFail(t *testing.T) {
	sut := makeListNextTaskUseCase()
	monkey.PatchInstanceMethod(
		reflect.TypeOf(sut.TaskRepository),
		"ListNextTasks",
		func(t *ListTaskRepositoryStub) ([]task.Task, error) {
			return nil, errors.New("error on load tasks")
		})
	_, err := sut.Execute()
	fmt.Println(err)
	assert.Error(t, err, "error on load task")
	monkey.UnpatchInstanceMethod(reflect.TypeOf(sut.TaskRepository), "ListNextTasks")
}
func TestListNextTaskUseCase(t *testing.T) {
	sut := makeListNextTaskUseCase()
	tasks, err := sut.Execute()
	assert.Nil(t, err)
	assert.Equal(t, len(tasks), 1)
	assert.Equal(t, tasks[0].Id, 1)
	assert.Equal(t, tasks[0].Description, "new task")
}

func TestListNextTaskUseCaseReturnsListOfReadTaskDTO(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2022, 11, 19, 13, 25, 0, 0, time.Local)
	})
	sut := makeListNextTaskUseCase()
	tasks, err := sut.Execute()
	assert.Nil(t, err)
	assert.Equal(t, len(tasks), 1)
	assert.Equal(t, tasks[0].Id, 1)
	assert.Equal(t, tasks[0].Age, "1 hour 25 minutes")
	assert.Equal(t, tasks[0].Priority, "high")
}
func TestListTasNextkUseCaseMapTasksToReadTaskDTO(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2022, 11, 19, 13, 25, 0, 0, time.Local)
	})
	taskList := makeTaskList()
	sut := makeListNextTaskUseCase()
	readTasks := sut.mapTaskToReadTaskDTO(taskList)

	assert.Equal(t, len(readTasks), len(taskList))
	assert.Equal(t, readTasks[0].Age, taskList[0].Age())
	assert.Equal(t, readTasks[0].Priority, taskList[0].PriorityToString())
}
