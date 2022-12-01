package task_dto

import (
	"testing"
	"time"
	task "todo/internal/task/model"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)
func makeAddTaskDto() AddTaskDTO {
	return AddTaskDTO{
		Description:"new task",
		Priority:"normal",
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
func TestListTasNextkUseCaseMapTasksToReadTaskDTO(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2022, 11, 19, 13, 25, 0, 0, time.Local)
	})
	taskList := makeTaskList()
	readTasks := MapTaskToReadTaskDTO(taskList)

	assert.Equal(t, len(readTasks), len(taskList))
	assert.Equal(t, readTasks[0].Age, taskList[0].Age())
	assert.Equal(t, readTasks[0].Priority, taskList[0].PriorityToString())

}
func TestMapAddTaskDtoToTask(t *testing.T) {
	addTaskDto := makeAddTaskDto()
	description, priority, err := MapAddTaskDtoToTask(addTaskDto)

	assert.Equal(t, description, addTaskDto.Description)
	assert.Equal(t, priority, task.Normal)
	assert.Nil(t, err)
}
func TestMapAddTaskDtoToTaskReturnsErrorWhenPriorityIsInvalid(t *testing.T) {
	addTaskDto := makeAddTaskDto()
	addTaskDto.Priority ="an invalid priority"
	description, priority, err := MapAddTaskDtoToTask(addTaskDto)

	assert.Empty(t, description)
	assert.Equal(t, priority, task.TaskPriority(-1))
	assert.Error(t, err, "priority must be low, normal or high")
}