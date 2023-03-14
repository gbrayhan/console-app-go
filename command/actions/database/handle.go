// Package database package database to implement actions on database
package database

import (
	"github.com/gbrayhan/console-app-go/command/actions/database/backup"
	"github.com/spf13/cobra"
)

func HandlerRun(cmd *cobra.Command, args []string) {
	switch cmd.Flag("action").Value.String() {
	case "backup":
		backup.Backup(args, cmd.Flag("type").Value.String())
	case "restore":
		println("Restore")
	case "compare":
		println("Compare")
	default:
		println("No action")
	}

}
