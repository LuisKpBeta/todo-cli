package comands

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
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
	completeTasl := usecases.NewCompleteTaskUseCase(repository, repository)
	return completeTasl
}
func runCompleteTask(dbConn *sql.DB) func(cmd *cobra.Command, args []string) {
	completeTaskUseCase := buildCompleteTask(dbConn)
	return func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			checkForError(err, args[0])
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
func checkForError(err error, value string){
	var msg string
	if errors.Is(err, strconv.ErrSyntax) {
		msg = fmt.Sprint("invalid value \"", value, "\" for a task id\n")
	} else {
		msg = err.Error()
	}
	fmt.Println(msg)
}
