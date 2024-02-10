/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package task

import (
	"encoding/csv"
	"flow/cli/utils"
	"log"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func mapToTableRow(row []string) table.Row {
	return table.Row{row[0], row[1], row[2]}
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the tasks registered",
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
