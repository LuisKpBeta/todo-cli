package usecases

import (
	task "todo/internal/task/model"
	dto "todo/internal/task/usecases/task_dto"
)

type ListTaskUseCase struct {
	TaskRepository task.ListTaskRepositoryInterface
}

func NewListTaskUseCase(taskRepository task.ListTaskRepositoryInterface) *ListTaskUseCase {
	return &ListTaskUseCase{
		TaskRepository: taskRepository,
	}
}

func (l *ListTaskUseCase) Execute(listAll bool) ([]dto.ReadTaskDTO, error) {
	tasks, err := l.TaskRepository.ListTasks(listAll)
	if err != nil {
		return nil, err
	}
	readTasks := dto.MapTaskToReadTaskDTO(tasks)
	return readTasks, nil
}
