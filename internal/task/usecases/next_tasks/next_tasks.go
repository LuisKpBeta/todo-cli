package usecases

import task "todo/internal/task/model"

type ReadTaskDTO struct {
	Id          int
	Description string
	Status      bool
	Priority    string
	Age         string
}

type ListNextTaskUseCase struct {
	TaskRepository task.ListNextTasksRepositoryInterface
}

func NewListNextTaskUseCase(taskRepository task.ListNextTasksRepositoryInterface) *ListNextTaskUseCase {
	return &ListNextTaskUseCase{
		TaskRepository: taskRepository,
	}
}

func (l *ListNextTaskUseCase) Execute() ([]ReadTaskDTO, error) {
	tasks, err := l.TaskRepository.ListNextTasks()
	if err != nil {
		return nil, err
	}
	readTasks:=l.mapTaskToReadTaskDTO(tasks)
	return readTasks, nil
}
func (l *ListNextTaskUseCase) mapTaskToReadTaskDTO(taskList []task.Task) []ReadTaskDTO {
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