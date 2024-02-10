/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package database

import (
	"flow/cli/utils"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var decoded bool
var hideCreds bool

func handleDataVisibility(database DatabaseConfig) DatabaseConfig {
	const HIDDEN = "********"

	if hideCreds {
		database.User = HIDDEN
		database.Pass = HIDDEN
		return database
	}

	if !decoded {
		database.Pass = HIDDEN
		return database
	}

	return database
}

func listDatabases(databases []DatabaseConfig) {
	headers := getDatabaseConfigTableHeaders()
	var databaseRows []table.Row

	for _, db := range databases {
		db := handleDataVisibility(db)
		databaseRows = append(databaseRows, db.mapToTableRow())
	}

	utils.PrintTable(headers, databaseRows)
}

func filterDatabases(query string, databases []DatabaseConfig) []DatabaseConfig {

	return databases
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the databases",
	Long:  `Lists the databases`,
	Run: func(cmd *cobra.Command, args []string) {
		databases := getDatabasesFromConfig()
		if len(args) < 1 {
			listDatabases(databases)
			return
		}

		// TODO: Add search option as first arg[0]
		if args[0] != "" {
			databases := filterDatabases(args[0], databases)
			listDatabases(databases)
		}
	},
}

func init() {
	listCmd.Flags().BoolVarP(&decoded, "decoded", "d", false, "Shows password with databases")
	listCmd.Flags().BoolVar(&hideCreds, "hide", false, "Hides Users and Passwords")
}
