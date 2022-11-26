package comands

import (
	"database/sql"
	"fmt"

	"github.com/spf13/cobra"
)

func DeleteTaskComand(dbConn *sql.DB) *cobra.Command {
	comand := &cobra.Command{
		Use:   "delete <id>",
		Short: "delete a task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("task %s\n deleted", args[0])
		},
	}
	comand.DisableFlagsInUseLine = true
	return comand
}
