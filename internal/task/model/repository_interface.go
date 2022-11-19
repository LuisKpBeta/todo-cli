package model

type TaskRepositoryInterface interface {
	AddTask(task *Task) (int, error)
}
