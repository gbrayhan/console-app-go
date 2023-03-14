package backup

func Backup(args []string, engineType string) {
	println("Backup")

	switch engineType {
	case "mysql":
		mysqlBackup(args)
	case "psql":
		psqlBackup(args)
	default:
		println("No engine")
	}
}
