package usecases

import (
	task "todo/internal/task/model"
	dto "todo/internal/task/usecases/task_dto"
)



type AddTaskUseCase struct {
	TaskRepository task.AddTaskRepositoryInterface
}

func NewAddTaskUseCase(taskRepository task.AddTaskRepositoryInterface) *AddTaskUseCase {
	return &AddTaskUseCase{
		TaskRepository: taskRepository,
	}
}

func (a *AddTaskUseCase) Execute(newTaskDTO dto.AddTaskDTO) (*task.Task, error) {
	var err error
	if(newTaskDTO.Priority==""){
		newTaskDTO.Priority = "normal"
	}
	description, priority, err :=dto.MapAddTaskDtoToTask(newTaskDTO)
	if err != nil {
		return nil, err
	}
	newTask, err := task.NewTask(description, priority)
	if err != nil {
		return nil, err
	}
	err = a.TaskRepository.AddTask(newTask)
	if err != nil {
		return nil, err
	}
	return newTask, nil
}
