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
	stmt, _ := suite.Db.Prepare("INSERT INTO tasks (description, status, priority, created) values(?,?,?,?)")
	res, _ := stmt.Exec(task.Description, 0, task.Priority, task.Created)
	id, _ := res.LastInsertId()
	task.Id = int(id)
}
func (suite *TaskRepositoryTestSuite) TearDownTest(suiteName, testName string) {
	suite.Db.Query("DELETE FROM tasks;")
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
