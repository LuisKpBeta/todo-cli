package common

import (
	"fmt"
	"todo/internal/task/usecases/task_dto"
)

func PrettyListPrinter(taskList []task_dto.ReadTaskDTO) {
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