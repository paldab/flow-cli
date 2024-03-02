package task

import (
	"io"
	"log"
	"os"
	"path"
	"strings"
)

func CreateBackup(){
	sourceFile := getTaskHistoryFile()
		parentDir := path.Dir(sourceFile.Name())
		backupPrefix := ".backup"
		backupFileName := strings.Split(TASKHISTORYFILE, ".")[0] + backupPrefix + ".csv"

		backupFile, err := os.Create(path.Join(parentDir, backupFileName))

		if err != nil {
			log.Fatal(err.Error())
		}

		_, err = io.Copy(backupFile, sourceFile)

		if err != nil {
			log.Fatal(err.Error())
		}

		println("Backup of tasks has been created!")
}
