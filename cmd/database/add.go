package database

import (
	"log"

	"github.com/flow-cli/internal/database"
	"github.com/spf13/cobra"
)

var hostInput string
var userInput string
var passInput string
var dbTypeInput string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a database to the config",
	Long: `Example usage 

database add example-prod --host=0.0.0.0 -u user1 -p pass1`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Missing database name. Could not find value to add to databases!")
		}

		database.AddDatabase(args[0], hostInput, userInput, passInput, dbTypeInput)
	},
}

func init() {
	addCmd.Flags().StringVar(&hostInput, "host", "", "Define the host of the database")
	addCmd.MarkFlagRequired("host")

	addCmd.Flags().StringVarP(&userInput, "user", "u", "", "Define a user for the database")
	addCmd.MarkFlagRequired("user")

	addCmd.Flags().StringVarP(&passInput, "password", "p", "", "Define a password for the database")

	addCmd.Flags().StringVarP(&dbTypeInput, "type", "t", "", "Type of database. MySQL or Postgres")
	addCmd.MarkFlagRequired("type")
}
