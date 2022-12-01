package usecases

import (
	"errors"
	"fmt"
	task "todo/internal/task/model"
)

type CompleteTaskUseCase struct {
	CompleteTaskRepository task.CompleteTaskRepositoryInterface
	FindTaskRepository     task.FindTaskRepositoryInterface
}

func NewCompleteTaskUseCase(completeTask task.CompleteTaskRepositoryInterface, findTaskById task.FindTaskRepositoryInterface) *CompleteTaskUseCase {
	return &CompleteTaskUseCase{
		CompleteTaskRepository: completeTask,
		FindTaskRepository:     findTaskById,
	}
}

func (a *CompleteTaskUseCase) Execute(taskId int) (*task.Task, error) {
	task, err := a.FindTaskRepository.FindById(taskId)
	if err != nil {
		return nil, err
	}
	if task == nil {
		err_msg := fmt.Sprint("tasks not found with id: ", taskId)
		return nil, errors.New(err_msg)
	}
	if(task.Status){
		err_msg := fmt.Sprint("tasks with id \"",taskId,"\" already completed")
		return nil, errors.New(err_msg)
	}
	task.CompleteTask()
	err = a.CompleteTaskRepository.CompleteTask(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}
