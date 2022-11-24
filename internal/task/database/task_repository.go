package database

import (
	"database/sql"
	"time"
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
func (t *TaskRepository) CompleteTask(task *task.Task) error {
	newStatus := 0
	if task.Status {
		newStatus = 1
	}
	stmt, err := t.Db.Prepare("UPDATE tasks set status=? where id=?")
	checkErr(err)
	_, err = stmt.Exec(newStatus, task.Id)
	return err
}
func (t *TaskRepository) FindById(taskId int) (*task.Task, error) {
	var id int
	var descption, priority string
	var status int
	var created time.Time
	err := t.Db.QueryRow("SELECT * FROM tasks where id = ?", taskId).
		Scan(&id, &descption, &status, &priority, &created)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	task := &task.Task{
		Id:          id,
		Description: descption,
		Status:      (status == 1),
		Priority:    task.TaskPriority(priority),
		Created:     created,
	}
	return task, nil
}

func (t *TaskRepository) DeleteById(taskId int) error {
	stmt, err := t.Db.Prepare("DELETE FROM tasks where id=?")
	checkErr(err)
	_, err = stmt.Exec(taskId)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
