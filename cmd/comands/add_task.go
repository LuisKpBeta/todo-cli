package comands

import (
	"database/sql"
	"fmt"

	"github.com/spf13/cobra"
)

func AddTaskComand(dbConn *sql.DB) *cobra.Command {
	comand := &cobra.Command{
		Use:   "add <description> [-p <priority>]",
		Short: "crete a new task with description and priority (low, normal or high)",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			priority, _ := cmd.Flags().GetString("priority")
			fmt.Printf("Description %s\n priority %s\n", args[0], priority)
		},
	}
	comand.DisableFlagsInUseLine = true
	comand.Flags().StringP("priority", "p", "", "define task priority")
	return comand
}
