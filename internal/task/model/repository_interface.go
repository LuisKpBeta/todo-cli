package model

type TaskRepositoryInterface interface {
	AddTask(task *Task) error
}
