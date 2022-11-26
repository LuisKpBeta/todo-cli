package task_dto

import (
	"errors"
	"strings"
	task "todo/internal/task/model"
)

type AddTaskDTO struct {
	Description string
	Priority    string
}

type ReadTaskDTO struct {
	Id          int
	Description string
	Status      bool
	Priority    string
	Age         string
}
func MapTaskToReadTaskDTO(taskList []task.Task) []ReadTaskDTO {
	var readTasks []ReadTaskDTO
	for _, task := range taskList {
		readTask := ReadTaskDTO{
			Id:          task.Id,
			Description: task.Description,
			Status:      task.Status,
			Priority:    task.PriorityToString(),
			Age:         task.Age(),
		}
		readTasks = append(readTasks, readTask)
	}
	return readTasks
}

func MapAddTaskDtoToTask(addTask AddTaskDTO) (string, task.TaskPriority, error){
	var priority int
	switch strings.ToLower(addTask.Priority){
	case "low":
		priority = 2
	case "normal":
		priority = 1
	case "high":
		priority = 0
	default:
		return "", -1, errors.New("priority must be low, normal or high")
	}
	return addTask.Description, task.TaskPriority(priority), nil
}