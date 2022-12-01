package comands

import (
	"database/sql"
	"fmt"
	common "todo/cmd/common"
	repo "todo/internal/task/database"
	usecases "todo/internal/task/usecases/complete_task"

	"github.com/spf13/cobra"
)

func CompleteTaskComand(dbConn *sql.DB) *cobra.Command {
	comand := &cobra.Command{
		Use:   "complete <id>",
		Short: "mask a task as done",
		Args:  cobra.ExactArgs(1),
		Run:   runCompleteTask(dbConn),
	}
	comand.DisableFlagsInUseLine = true
	return comand
}
func buildCompleteTask(dbConn *sql.DB) *usecases.CompleteTaskUseCase {
	repository := repo.NewTaskRepository(dbConn)
	completeTask := usecases.NewCompleteTaskUseCase(repository, repository)
	return completeTask
}
func runCompleteTask(dbConn *sql.DB) func(cmd *cobra.Command, args []string) {
	completeTaskUseCase := buildCompleteTask(dbConn)
	return func(cmd *cobra.Command, args []string) {
		taskId, err := common.CheckAndParseIdArg(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		task, err := completeTaskUseCase.Execute(int(taskId))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("task:", task.Id, "completed in ", task.Age())
	}
}

