package database

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func Handler(cmd *cobra.Command, args []string) {
	fmt.Println("Hello from subcommand!")
	fmt.Println("Print: " + strings.Join(args, "|"))

}
