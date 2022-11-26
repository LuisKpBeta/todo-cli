package comands

import (
	"database/sql"
	"fmt"
	"strings"

	repo "todo/internal/task/database"
	usecases "todo/internal/task/usecases/add_task"
	dto "todo/internal/task/usecases/task_dto"

	"github.com/spf13/cobra"
)

func AddTaskComand(dbConn *sql.DB) *cobra.Command {
	addTaskUseCase := buildAddTask(dbConn)
	comand := &cobra.Command{
		Use:   "add <description> [-p <priority>]",
		Short: "crete a new task with description and priority (low, normal or high)",
		Run: func(cmd *cobra.Command, args []string) {
			priority, _ := cmd.Flags().GetString("priority")
			addTaskDto := dto.AddTaskDTO{
				Description: strings.Join(args, " "),
				Priority:    priority,
			}
			task, err := addTaskUseCase.Execute(addTaskDto)
			if err != nil {
				fmt.Println("WARNING")
				fmt.Println(err)
			}else {
				fmt.Println("TASK \"",task.Description,"\" criada com sucesso")
			}
		},
	}
	comand.DisableFlagsInUseLine = true
	comand.Flags().StringP("priority", "p", "", "define task priority")
	comand.MarkFlagRequired("priority")
	return comand
}
func buildAddTask(dbConn *sql.DB) *usecases.AddTaskUseCase {
	repository := repo.NewTaskRepository(dbConn)
	addTask := usecases.NewAddTaskUseCase(repository)
	return addTask
}
