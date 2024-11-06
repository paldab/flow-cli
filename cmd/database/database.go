package database

import (
	"github.com/spf13/cobra"
)

var DatabaseCmd = &cobra.Command{
	Use:   "database",
    Aliases: []string{"db"},
	Short: "database represents the database module of the CLI. Managing your databases for easy overview and connection",
	Long:  `database represents the database module of the CLI. Managing your databases for easy overview and connection`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	DatabaseCmd.AddCommand(addCmd)
	DatabaseCmd.AddCommand(listCmd)
	DatabaseCmd.AddCommand(deleteCmd)
	DatabaseCmd.AddCommand(whitelistCmd)
	DatabaseCmd.AddCommand(connectCmd)
}
