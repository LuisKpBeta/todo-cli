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
type DeleteTaskRepositoryInterface interface {
	DeleteById(taskId int) error
}
type ListTaskRepositoryInterface interface {
	ListTasks(listAll bool) ([]Task, error)
}
