package model

type AddTaskRepositoryInterface interface {
	AddTask(task *Task) error
}
type CompleteTaskRepositoryInterface interface {
	CompleteTask(task *Task) error
}
type FindTaskRepositoryInterface interface {
	FindById(taskId int) (*Task, error)
}
