package comands

import (
	"database/sql"
	"fmt"
	common "todo/cmd/common"
	repo "todo/internal/task/database"
	usecases "todo/internal/task/usecases/delete_task"

	"github.com/spf13/cobra"
)

func DeleteTaskComand(dbConn *sql.DB) *cobra.Command {
	comand := &cobra.Command{
		Use:   "delete <id>",
		Short: "delete a task",
		Args:  cobra.ExactArgs(1),
		Run: runDeleteTask(dbConn),
	}
	comand.DisableFlagsInUseLine = true
	return comand
}

func buildDeleteTask(dbConn *sql.DB) *usecases.DeleteTaskUseCase {
	repository := repo.NewTaskRepository(dbConn)
	completeTasl := usecases.NewDeleteTaskUseCase(repository)
	return completeTasl
}
func runDeleteTask(dbConn *sql.DB) func(cmd *cobra.Command, args []string) {
	deleteTaskUseCase := buildDeleteTask(dbConn)
	return func(cmd *cobra.Command, args []string) {
		taskId, err := common.CheckAndParseIdArg(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		err = deleteTaskUseCase.Execute(int(taskId))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("deletion succeed")
	}
}