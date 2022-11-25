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
	_, err := t.Db.Exec("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, description VARCHAR(256) NULL,	status tinyint default 0,	priority tinyint,created DATE NULL);")
	checkErr(err)
}
func (t *TaskRepository) AddTask(task *task.Task) error {
	stmt, err := t.Db.Prepare("INSERT INTO tasks (description, status, priority, created) values(?,?,?,?)")
	checkErr(err)

	status := 0
	if task.Status {
		status = 1
	}
	result, err := stmt.Exec(task.Description, status, task.Priority(), task.Created)
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
	var descption string
	var status, priority int
	var created time.Time
	err := t.Db.QueryRow("SELECT * FROM tasks where id = ?", taskId).
		Scan(&id, &descption, &status, &priority, &created)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	tsk := &task.Task{
		Id:          id,
		Description: descption,
		Status:      (status == 1),
		Created:     created,
	}
	tsk.SetPriority(task.TaskPriority(priority))
	return tsk, nil
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
func (t *TaskRepository) ListTasks(listAll bool, orderByPriority bool) ([]task.Task, error) {
	query := "SELECT * FROM tasks"
	if !listAll {
		query = query + " WHERE status=0"
	}
	if orderByPriority {
		query = query + " ORDER BY priority"
	}
	rows, err := t.Db.Query(query)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	var taskList []task.Task
	for rows.Next() {
		tsk := task.Task{}
		var status, priority int
		err := rows.Scan(&tsk.Id, &tsk.Description, &status, &priority, &tsk.Created)
		if err != nil {
			return nil, err
		}
		tsk.Status = status == 1
		tsk.SetPriority(task.TaskPriority(priority))
		taskList = append(taskList, tsk)
	}

	return taskList, nil
}
func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
