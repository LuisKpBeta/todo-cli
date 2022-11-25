package usecases

import task "todo/internal/task/model"

type ListTaskArgs struct {
	ListAll         bool
	OrderByPriority bool
}

type ListTaskUseCase struct {
	TaskRepository task.ListTaskRepositoryInterface
}

func NewListTaskUseCase(taskRepository task.ListTaskRepositoryInterface) *ListTaskUseCase {
	return &ListTaskUseCase{
		TaskRepository: taskRepository,
	}
}

func (l *ListTaskUseCase) Execute(listTaskArgs ListTaskArgs) ([]task.Task, error) {
	tasks, err := l.TaskRepository.ListTasks(listTaskArgs.ListAll, listTaskArgs.OrderByPriority)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
