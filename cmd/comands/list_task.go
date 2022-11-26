package comands

import (
	"database/sql"
	"fmt"
	repo "todo/internal/task/database"
	usecases "todo/internal/task/usecases/list_tasks"

	"github.com/spf13/cobra"
)

func ListTasksComand(dbConn *sql.DB) *cobra.Command {
	listTasksUseCase := buildLisTask(dbConn)
	comand := &cobra.Command{
		Use:   "list [-a]",
		Short: "show pending tasks. use -a (--all) for all tasks",
		Run: func(cmd *cobra.Command, _ []string) {
			taskList, _ := listTasksUseCase.Execute(false)
			fmt.Println(taskList)
		},
	}
	comand.DisableFlagsInUseLine = true
	return comand
}

func buildLisTask(dbConn *sql.DB) *usecases.ListTaskUseCase {
	repository := repo.NewTaskRepository(dbConn)
	listTask := usecases.NewListTaskUseCase(repository)
	return listTask
}
