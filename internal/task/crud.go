package task

import (
	"encoding/csv"
	"log"
	"strings"
	"time"

	"github.com/flow-cli/internal/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func AddTask(name, taskType string) {
		if taskType == "" {
			keywords := []string{"fix", "bugfix"}

			for _, keyword := range keywords {
				if strings.Contains(strings.ToLower(name), keyword) {
					taskType = "Bugfix"
					break
				}
			}

			if taskType == "" {
				taskType = "Feature"
			}
		}

		currentTime := time.Now()
		currentDate := currentTime.Format("2006-01-02")
		var task Task = Task{
			Name: name,
			Type: cases.Title(language.English, cases.Compact).String(strings.ToLower(taskType)),
			Date: currentDate,
		}

		writeToCsv(getTaskHistoryFile(), task.toRecord())
}

func ListTasks(){
	csvReader := csv.NewReader(getTaskHistoryFile())
	var headers table.Row
	var tasks []table.Row

	rows, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err.Error())
	}

	for index, row := range rows {
		if index == 0 {
			headers = mapToTableRow(row)
			continue
		}
		tasks = append(tasks, mapToTableRow(row))
	}

	utils.PrintTable(headers, tasks)
}