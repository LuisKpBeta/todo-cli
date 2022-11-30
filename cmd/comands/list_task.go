package comands

import (
	"database/sql"
	"fmt"
	repo "todo/internal/task/database"
	usecases "todo/internal/task/usecases/list_tasks"
	"todo/internal/task/usecases/task_dto"

	"github.com/spf13/cobra"
)

func ListTasksComand(dbConn *sql.DB) *cobra.Command {
	comand := &cobra.Command{
		Use:   "list [-a]",
		Short: "show pending tasks. use -a (--all) for all tasks",
		Run:   runListTask(dbConn),
	}
	comand.Flags().BoolP("all", "a", false, "list all tasks")
	comand.DisableFlagsInUseLine = true
	return comand
}

func buildLisTask(dbConn *sql.DB) *usecases.ListTaskUseCase {
	repository := repo.NewTaskRepository(dbConn)
	listTask := usecases.NewListTaskUseCase(repository)
	return listTask
}
func runListTask(dbConn *sql.DB) func(cmd *cobra.Command, args []string) {
	listTasksUseCase := buildLisTask(dbConn)

	return func(cmd *cobra.Command, _ []string) {
		getAll, _ := cmd.Flags().GetBool("all")
		taskList, _ := listTasksUseCase.Execute(getAll)
		prettyListPrinter(taskList)
	}
}
func prettyListPrinter(taskList []task_dto.ReadTaskDTO) {
	for _, task := range taskList {
		fmt.Printf("%d. %s %s (%s) %s\n", task.Id, statusIcon(task.Status), task.Description, task.Priority, task.Age)
	}
}
func statusIcon(status bool) string {
	if status {
		return "✔"
	}
	return "□"
}
