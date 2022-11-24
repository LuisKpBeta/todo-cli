package usecases

import (
	task "todo/internal/task/model"
)

type DeleteTaskUseCase struct {
	DeleteTaskRepository task.DeleteTaskRepositoryInterface
}

func NewDeleteTaskUseCase(deleteTask task.DeleteTaskRepositoryInterface) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{
		DeleteTaskRepository: deleteTask,
	}
}

func (d *DeleteTaskUseCase) Execute(taskId int) error {
	err := d.DeleteTaskRepository.DeleteById(taskId)
	return err
}
