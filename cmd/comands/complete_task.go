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
	completeTaskUseCase := buildCompleteTask(dbConn)
	comand := &cobra.Command{
		Use:   "complete <id>",
		Short: "mask a task as done",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			taskId, err :=  strconv.Atoi(args[0])
			if(err!=nil){
				if errors.Is(err, strconv.ErrSyntax) {
					fmt.Printf("invalid value \"%s\" for a task id\n",args[0])
				}else{
					fmt.Println(err.Error())
				}
			}else{
				task, err:= completeTaskUseCase.Execute(int(taskId))
				if(err!=nil){
					fmt.Println(err.Error())
					return
				}
				fmt.Println("task:", task.Id, "completed in ", task.Age())
			}
		},
	}
	comand.DisableFlagsInUseLine = true
	return comand
}
func buildCompleteTask(dbConn *sql.DB) *usecases.CompleteTaskUseCase {
	repository := repo.NewTaskRepository(dbConn)
	completeTasl := usecases.NewCompleteTaskUseCase(repository, repository)
	return completeTasl
}