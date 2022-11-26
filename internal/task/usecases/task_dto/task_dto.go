package task_dto

import task "todo/internal/task/model"

type AddTaskDTO struct {
	Description string
	Priority    task.TaskPriority
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