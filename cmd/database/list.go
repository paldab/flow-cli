package database

import (
	"github.com/flow-cli/internal/database"
	"github.com/spf13/cobra"
)

var decoded bool
var hideCreds bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the databases",
	Long:  `Lists the databases. First argument is used to search within the list of databases`,
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		databases := database.GetDatabasesFromConfig()
		database.ListDatabases(query, hideCreds, decoded, databases)
	},
}

func init() {
	listCmd.Flags().BoolVarP(&decoded, "decoded", "d", false, "Shows password with databases")
	listCmd.Flags().BoolVar(&hideCreds, "hide", false, "Hides Users and Passwords")
}
