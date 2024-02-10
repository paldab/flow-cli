/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package task

import (
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task",

	Run: func(cmd *cobra.Command, args []string) {

		taskName, err := cmd.Flags().GetString("task")
		if err != nil {
			log.Fatal(err.Error())
		}

		taskType, err := cmd.Flags().GetString("type")
		if err != nil {
			log.Fatal(err.Error())
		}

		if taskType == "" {
			keywords := []string{"fix", "bugfix"}

			for _, keyword := range keywords {
				if strings.Contains(strings.ToLower(taskName), keyword) {
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
			Name: taskName,
			Type: cases.Title(language.English, cases.Compact).String(strings.ToLower(taskType)),
			Date: currentDate,
		}

		writeToCsv(getTaskHistoryFile(), task.toRecord())
	},
}

func init() {
	addCmd.Flags().String("task", "", "Name of the task")
	addCmd.Flags().String("type", "", "Type of task. Feature, Bugfix, etc...")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
