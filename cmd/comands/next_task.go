package comands

import (
	"database/sql"
	"fmt"

	"github.com/spf13/cobra"
)

func NextTasksComand(dbConn *sql.DB) *cobra.Command {
	comand := &cobra.Command{
		Use:   "next",
		Short: "show next task of each priority",
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Printf("next tasks")
		},
	}
	comand.DisableFlagsInUseLine = true
	return comand
}

