package backup

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func psqlBackup(args []string) (err error) {
	pwd := "Password1"
	err = os.Setenv("PGPASSWORD", pwd)
	if err != nil {
		fmt.Println(err)
		return err
	}

	dbList := []string{
		"aislelinx", "archiveandpurging", "asrsmanager", "businessorchestrator",
		"capacitymanager", "clientservices", "container", "crane", "grafanadb",
		"inventory", "keycloak", "liftlinx", "light", "location", "momentumconnect",
		"momentumrdb", "notificationmanager", "order", "postgres", "print", "putaway",
		"robotics", "routing", "routingclient", "shuttle", "smartinventoryallocation",
		"smartrelease", "socketproxy", "subscription", "taskmanager", "traveltimeprediction",
		"wave", "workassignment",
	}

	// Get the current local date and time
	dt := time.Now()
	datestamp := dt.Format("2006-01-02")
	timestamp := dt.Format("15-04-05")

	// Loop through the database list and dump each database
	for _, dbName := range dbList {
		// Construct the dump file path
		dumpFilePath := fmt.Sprintf("C:\\app\\database\\backup\\clientservices-%s-%s_%s-dump.sql", dbName, datestamp, timestamp)

		// Construct the pg_dump command
		cmd := exec.Command("C:/Program Files/PostgreSQL/12/bin/pg_dump.exe",
			"--clean", "--if-exists", fmt.Sprintf("--file=%s", dumpFilePath),
			"--column-inserts", "--rows-per-insert=10000",
			"--dbname="+dbName, "--username=sa",
			"--host=10.12.183.125", "--port=8432")
		cmd.Env = append(os.Environ(), "PGPASSWORD=Password1")

		// Run the pg_dump command
		if err := cmd.Run(); err != nil {
			fmt.Printf("Failed to dump database %s: %s\n", dbName, err.Error())
		} else {
			fmt.Printf("Successfully dumped database %s to %s\n", dbName, dumpFilePath)
		}
	}
	return
}
