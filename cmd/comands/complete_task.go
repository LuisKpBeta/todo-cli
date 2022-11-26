package comands

import (
	"database/sql"
	"fmt"

	"github.com/spf13/cobra"
)

func CompleteTaskComand(dbConn *sql.DB) *cobra.Command {
	comand := &cobra.Command{
		Use:   "complete <id>",
		Short: "mask a task as done",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("task %s\n finished", args[0])
		},
	}
	comand.DisableFlagsInUseLine = true
	return comand
}
