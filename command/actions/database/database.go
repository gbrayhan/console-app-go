// Package database package database to implement actions on database
package database

import (
	"github.com/spf13/cobra"
)

func Action() *cobra.Command {
	subCommand := cobra.Command{
		Use: "database",
		Run: HandlerRun,
	}
	var typeFlag string
	var actionFlag string
	subCommand.PersistentFlags().StringVar(&typeFlag, "type", "sql", "type of engine database (psql, mysql, sql, etc)")
	subCommand.PersistentFlags().StringVar(&actionFlag, "action", "backup", "action to do (backup, restore, etc)")
	return &subCommand
}
