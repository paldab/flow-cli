package task

import (
	"encoding/csv"
	"flow/cli/utils"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const TASKHISTORYFILE = "flow_taskhistory.csv"

type Task struct {
	Name string
	Type string
	Date string
}

func (t Task) toRecord() []string {
	return []string{t.Name, t.Type, t.Date}
}

func getTaskHistoryFile() *os.File {
	taskHistoryPath := viper.GetString("tasks.history.path")

	if taskHistoryPath == "" {
		log.Fatal("taskHistoryPath is an empty string")
	}

	file, err := os.OpenFile(taskHistoryPath, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func writeToCsv(file *os.File, record []string) {
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write(record)
}

func SetTaskHeaders() {
	file := getTaskHistoryFile()
	csvReader := csv.NewReader(file)

	_, err := csvReader.Read()

	if err != nil {
		if err != io.EOF {
			log.Fatal(err.Error())
		}

		taskHeaders := utils.GetStructKeys(Task{})
		writeToCsv(file, taskHeaders)
	}
}

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Task keeps track if the tasks that you have done. Use one of it's subfunctions!",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	TaskCmd.AddCommand(addCmd)
	TaskCmd.AddCommand(listCmd)
	TaskCmd.AddCommand(deleteCmd)
	TaskCmd.AddCommand(createBackupCmd)
}
