package database

import (
	"database/sql"
	task "todo/internal/task/model"
)

type TaskRepository struct {
	Db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{Db: db}
}
func (t *TaskRepository) CreateTaskTableIfNoExists() {
	_, err := t.Db.Exec("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, description VARCHAR(256) NULL,	status tinyint default 0,	priority VARCHAR(7),created DATE NULL);")
	checkErr(err)
}
func (t *TaskRepository) AddTask(task *task.Task) error {
	stmt, err := t.Db.Prepare("INSERT INTO tasks (description, status, priority, created) values(?,?,?,?)")
	checkErr(err)

	status := 0
	if task.Status {
		status = 1
	}
	result, err := stmt.Exec(task.Description, status, task.Priority, task.Created)
	if err != nil {
		return err
	}
	createdId, err := result.LastInsertId()
	checkErr(err)
	task.Id = int(createdId)
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
