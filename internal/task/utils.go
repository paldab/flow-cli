package task

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/flow-cli/internal/utils"
	"github.com/jedib0t/go-pretty/v6/table"
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

func mapToTableRow(row []string) table.Row {
	return table.Row{row[0], row[1], row[2]}
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