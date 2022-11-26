package usecases

import (
	task "todo/internal/task/model"
	dto "todo/internal/task/usecases/task_dto"
)

type ListNextTaskUseCase struct {
	TaskRepository task.ListNextTasksRepositoryInterface
}

func NewListNextTaskUseCase(taskRepository task.ListNextTasksRepositoryInterface) *ListNextTaskUseCase {
	return &ListNextTaskUseCase{
		TaskRepository: taskRepository,
	}
}

func (l *ListNextTaskUseCase) Execute() ([]dto.ReadTaskDTO, error) {
	tasks, err := l.TaskRepository.ListNextTasks()
	if err != nil {
		return nil, err
	}
	readTasks := dto.MapTaskToReadTaskDTO(tasks)
	return readTasks, nil
}

