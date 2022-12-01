package comands

import (
	"database/sql"
	c "todo/cmd/common"
	repo "todo/internal/task/database"
	usecases "todo/internal/task/usecases/next_tasks"

	"github.com/spf13/cobra"
)

func NextTasksComand(dbConn *sql.DB) *cobra.Command {
	comand := &cobra.Command{
		Use:   "next",
		Short: "show next task of each priority",
		Run: runListNextTask(dbConn),
	}
	comand.DisableFlagsInUseLine = true
	return comand
}

func buildLisNextTask(dbConn *sql.DB) *usecases.ListNextTaskUseCase {
	repository := repo.NewTaskRepository(dbConn)
	listTask := usecases.NewListNextTaskUseCase(repository)
	return listTask
}
func runListNextTask(dbConn *sql.DB) func(cmd *cobra.Command, args []string) {
	listTasksUseCase := buildLisNextTask(dbConn)

	return func(cmd *cobra.Command, _ []string) {
		taskList, _ := listTasksUseCase.Execute()
		c.PrettyListPrinter(taskList)
	}
}

