package backup

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func mysqlBackup(args []string) {
	mapArgs := make(map[string]string)
	for _, arg := range args {
		parts := strings.Split(arg, "=")
		mapArgs[parts[0]] = parts[1]
	}

	err := validateMapArgs(mapArgs)
	if err != nil {
		println(err.Error())
		return
	}
	println("Validated mapArgs")

	timestamp := time.Now().Format("2006-01-02-15-04-05")
	tables := strings.Split(mapArgs["tables"], ",")

	if len(tables) > 0 && tables[0] != "" {
		println("Dumping tables")
		for _, table := range tables {
			cmd := exec.Command("mysqldump", "--host="+mapArgs["db_host"], "--port="+mapArgs["db_port"], "--user="+mapArgs["db_user"], "--password="+mapArgs["db_password"],
				"--single-transaction", "--routines", "--triggers", "--add-drop-database", "--column-statistics=0", mapArgs["db_name"], table)
			output, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
				return
			}

			filename := fmt.Sprintf(mapArgs["dir_backup"]+"/%s-%s-%s-dump.sql", timestamp, mapArgs["db_name"], table)
			err = os.WriteFile(filename, output, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		return
	}

	println("No tables specified, dumping entire database")
	cmd := exec.Command("mysqldump", "--host="+mapArgs["db_host"], "--port="+mapArgs["db_port"], "--user="+mapArgs["db_user"], "--password="+mapArgs["db_password"],
		"--single-transaction", "--routines", "--triggers", "--add-drop-database", "--column-statistics=0", mapArgs["db_name"])
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("some error:", err)
		return
	}

	filename := fmt.Sprintf(mapArgs["dir_backup"]+"/%s-%s-dump.sql", timestamp, mapArgs["db_name"])
	err = os.WriteFile(filename, output, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func validateMapArgs(mapArgs map[string]string) error {
	requiredArgs := []string{"db_host", "db_port", "db_user", "db_password", "db_name", "dir_backup"}
	missingArgs := make([]string, 0)
	for _, arg := range requiredArgs {
		if _, ok := mapArgs[arg]; !ok {
			missingArgs = append(missingArgs, arg)
		}
	}
	if len(missingArgs) > 0 {
		return fmt.Errorf("missing arguments: %s", strings.Join(missingArgs, ", "))
	}
	return nil
}
