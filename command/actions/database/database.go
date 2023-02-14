package database

import (
	"github.com/spf13/cobra"
)

func Action() *cobra.Command {
	subCommand := cobra.Command{
		Use: "database",
		Run: Handler,
	}
	var exampleFlag string
	subCommand.PersistentFlags().StringVar(&exampleFlag, "example", "", "an example flag with a value")

	subCommand.SetArgs([]string{"action", "arg1", "arg2"})

	return &subCommand
}
