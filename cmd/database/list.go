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
	Long:  `Lists the databases`,
	Run: func(cmd *cobra.Command, args []string) {
		databases := database.GetDatabasesFromConfig()
		if len(args) < 1 {
			database.ListDatabases(hideCreds, decoded, databases)
			return
		}

		// TODO: Add search option as first arg[0]
		if args[0] != "" {
			databases := database.FilterDatabases(args[0], databases)
			database.ListDatabases(hideCreds, decoded, databases)
		}
	},
}

func init() {
	listCmd.Flags().BoolVarP(&decoded, "decoded", "d", false, "Shows password with databases")
	listCmd.Flags().BoolVar(&hideCreds, "hide", false, "Hides Users and Passwords")
}
