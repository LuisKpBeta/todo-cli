package database

import (
	"database/sql"
	"testing"
	"time"
	task "todo/internal/task/model"

	"bou.ke/monkey"
	_ "github.com/mattn/go-sqlite3"

	"github.com/stretchr/testify/suite"
)

type TaskRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *TaskRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	repo := NewTaskRepository(db)
	repo.CreateTaskTableIfNoExists()
	suite.Db = db
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2022, 11, 19, 12, 59, 59, 0, time.UTC)
	})
}
func (suite *TaskRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}
func (suite *TaskRepositoryTestSuite) insertDummyTask(task *task.Task) {
	status := 0
	if task.Status {
		status = 1
	}
	stmt, _ := suite.Db.Prepare("INSERT INTO tasks (description, status, priority, created) values(?,?,?,?)")
	res, _ := stmt.Exec(task.Description, status, task.Priority(), task.Created)
	id, _ := res.LastInsertId()
	task.Id = int(id)
}
func (suite *TaskRepositoryTestSuite) SetupTest() {
	stmt, err := suite.Db.Prepare("DELETE FROM tasks")
	checkErr(err)
	stmt.Exec()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}

func (suite *TaskRepositoryTestSuite) TestReturnTaskIdOnCreateTask() {
	newTask, err := task.NewTask("nova task", task.Low)
	suite.NoError(err)
	repo := NewTaskRepository(suite.Db)
	err = repo.AddTask(newTask)
	suite.NoError(err)

	var description string
	var id int
	err = suite.Db.QueryRow("Select id, description from tasks where id = ?", newTask.Id).
		Scan(&id, &description)

	suite.NoError(err)
	suite.Equal(newTask.Description, description)
	suite.Equal(newTask.Id, id)
}

func (suite *TaskRepositoryTestSuite) TestReturnTaskOnFindById() {
	newTask, err := task.NewTask("nova task", task.Low)
	suite.NoError(err)
	suite.insertDummyTask(newTask)
	repo := NewTaskRepository(suite.Db)
	task, err := repo.FindById(newTask.Id)
	suite.NoError(err)

	suite.Equal(newTask.Description, task.Description)
	suite.Equal(newTask.Id, task.Id)
	suite.Equal(newTask.Status, task.Status)
	suite.Equal(newTask.Created, task.Created)
}
func (suite *TaskRepositoryTestSuite) TestFindTaskByIdWhenTaskNoExists() {
	newTask, err := task.NewTask("nova task", task.Low)
	suite.NoError(err)
	repo := NewTaskRepository(suite.Db)
	task, err := repo.FindById(newTask.Id)

	suite.Nil(task)
	suite.Nil(err)
}
func (suite *TaskRepositoryTestSuite) TestCompleteTaskRepository() {
	newTask, err := task.NewTask("nova task", task.Low)
	suite.NoError(err)
	suite.insertDummyTask(newTask)
	newTask.CompleteTask()
	repo := NewTaskRepository(suite.Db)
	err = repo.CompleteTask(newTask)
	suite.NoError(err)

	var status int
	var id int
	err = suite.Db.QueryRow("Select id, status from tasks where id = ?", newTask.Id).
		Scan(&id, &status)

	suite.Nil(err)
	suite.Equal(newTask.Id, id)
	suite.Equal(newTask.Status, (status == 1))
}
func (suite *TaskRepositoryTestSuite) TestDeleteTaskRepositoryDeleteTaskById() {
	newTask, err := task.NewTask("nova task", task.Low)
	suite.NoError(err)
	suite.insertDummyTask(newTask)
	repo := NewTaskRepository(suite.Db)
	err = repo.DeleteById(newTask.Id)
	suite.Nil(err)

	var id int
	suite.Db.QueryRow("Select id from tasks where id = ?", newTask.Id).
		Scan(&id)

	suite.Zero(id)
}
func (suite *TaskRepositoryTestSuite) TestListTaskRepository() {
	task1, _ := task.NewTask("nova 1", task.Low)
	completedTask, _ := task.NewTask("task 2", task.Normal)
	completedTask.Status = true
	suite.insertDummyTask(task1)
	suite.insertDummyTask(completedTask)
	repo := NewTaskRepository(suite.Db)
	taskList, err := repo.ListTasks(true)
	suite.Nil(err)

	suite.Equal(len(taskList), 2)
	suite.Equal(taskList[0].Id, task1.Id)
	suite.Equal(taskList[1].Id, completedTask.Id)
}
func (suite *TaskRepositoryTestSuite) TestListTaskRepositoryGetOnlyPendingTasks() {
	task1, _ := task.NewTask("task 1", task.Low)
	completedTask, _ := task.NewTask("task 2", task.Normal)
	completedTask.Status = true
	suite.insertDummyTask(task1)
	suite.insertDummyTask(completedTask)
	repo := NewTaskRepository(suite.Db)
	taskList, err := repo.ListTasks(false)
	suite.Nil(err)
	for _, task := range taskList {
		suite.False(task.Status)
	}
}
func (suite *TaskRepositoryTestSuite) TestListNextTaskRepository() {
	monkey.Unpatch(time.Now)
	createTask := func(desc string, priority task.TaskPriority, created time.Time, status bool) *task.Task {
		tsk, _ := task.NewTask(desc, priority)
		tsk.Created = created
		tsk.Status = status
		suite.insertDummyTask(tsk)
		return tsk
	}
	oldTime := time.Date(2022, 11, 19, 12, 0, 0, 0, time.UTC)
	newTime := time.Date(2022, 11, 20, 12, 0, 0, 0, time.UTC)
	createTask("normal 1", task.Normal, oldTime, false)
	createTask("high 1", task.High, oldTime, true)
	createTask("low 1", task.Low, oldTime, false)
	createTask("normal 2", task.Normal, newTime, false)
	createTask("low 2", task.Low, newTime, false)
	createTask("high 2", task.High, newTime, false)

	repo := NewTaskRepository(suite.Db)
	taskList, err := repo.ListNextTasks()
	suite.Nil(err)
	suite.Equal(len(taskList), 3)
	suite.Equal(taskList[0].Description, "high 2")
	suite.Equal(taskList[0].Priority(), task.High)
	suite.Equal(taskList[1].Description, "normal 1")
	suite.Equal(taskList[1].Priority(), task.Normal)
	suite.Equal(taskList[2].Description, "low 1")
	suite.Equal(taskList[2].Priority(), task.Low)
}
