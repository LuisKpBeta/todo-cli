package usecases

import task "todo/internal/task/model"

type ListTaskArgs struct {
	ListAll         bool
	OrderByPriority bool
}
type ReadTaskDTO struct {
	Id          int
	Description string
	Status      bool
	Priority    string
	Age         string
}

type ListTaskUseCase struct {
	TaskRepository task.ListTaskRepositoryInterface
}

func NewListTaskUseCase(taskRepository task.ListTaskRepositoryInterface) *ListTaskUseCase {
	return &ListTaskUseCase{
		TaskRepository: taskRepository,
	}
}

func (l *ListTaskUseCase) Execute(listTaskArgs ListTaskArgs) ([]ReadTaskDTO, error) {
	tasks, err := l.TaskRepository.ListTasks(listTaskArgs.ListAll, listTaskArgs.OrderByPriority)
	if err != nil {
		return nil, err
	}
	readTasks:=l.mapTaskToReadTaskDTO(tasks)
	return readTasks, nil
}
func (l *ListTaskUseCase) mapTaskToReadTaskDTO(taskList []task.Task) []ReadTaskDTO {
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