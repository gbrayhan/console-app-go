package main

import (
	"fmt"
	"github.com/gbrayhan/console-app-go/command/actions/database"
	"github.com/mbndr/figlet4go"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "Automated Tools",
		Run: func(cmd *cobra.Command, args []string) {
			ascii := figlet4go.NewAsciiRender()
			renderStr, _ := ascii.Render("BossonH")
			fmt.Print(renderStr)
		},
	}

	rootCmd.AddCommand(database.Action())
	rootCmd.Execute()
}
