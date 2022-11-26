package comands

import (
	"database/sql"
	"fmt"

	"github.com/spf13/cobra"
)

func ListTasksComand(dbConn *sql.DB) *cobra.Command {
	comand := &cobra.Command{
		Use:   "list [-a]",
		Short: "show pending tasks. use -a (--all) for all tasks",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("task %s\n finished", args[0])
		},
	}
	comand.DisableFlagsInUseLine = true
	return comand
}

