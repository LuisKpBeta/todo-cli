package database

import (
	"database/sql"
	"testing"
	task "todo/internal/task/model"

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
}
func (suite *TaskRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func (suite *TaskRepositoryTestSuite) AfterTest(suiteName, testName string) {
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
