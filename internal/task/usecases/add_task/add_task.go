package usecases

import task "todo/internal/task/model"

type AddTaskDTO struct {
	Description string
	Priority    task.TaskPriority
}

type AddTaskUseCase struct {
	TaskRepository task.AddTaskRepositoryInterface
}

func NewAddTaskUseCase(taskRepository task.AddTaskRepositoryInterface) *AddTaskUseCase {
	return &AddTaskUseCase{
		TaskRepository: taskRepository,
	}
}

func (a *AddTaskUseCase) Execute(newTaskDTO AddTaskDTO) (*task.Task, error) {
	var err error
	newTask, err := task.NewTask(newTaskDTO.Description, newTaskDTO.Priority)
	if err != nil {
		return nil, err
	}
	err = a.TaskRepository.AddTask(newTask)
	if err != nil {
		return nil, err
	}
	return newTask, nil
}
