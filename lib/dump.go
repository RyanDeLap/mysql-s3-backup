package lib

import (
	"fmt"
	"os"
	"os/exec"
)

func DumpFiles(fileName string) error {
	fmt.Println("Beginning dump of MySQL Database, this can take a while.")

	sqlUsername := os.Getenv("MYSQL_USERNAME")
	sqlPassword := os.Getenv("MYSQL_PASSWORD")
	sqlDatabase := os.Getenv("MYSQL_DATABASE")
	
	dumpCommand := fmt.Sprintf("mysqldump --user=%s --password=%s %s | gzip -9 > %s", sqlUsername, sqlPassword, sqlDatabase, fileName)

	cmd := exec.Command("bash", "-c", dumpCommand)
	err := cmd.Run()

	return err
}
