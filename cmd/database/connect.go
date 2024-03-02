package database

import (
	"log"

	"github.com/flow-cli/internal/database"
	"github.com/spf13/cobra"
)

var passwordInput string
var databaseInput string

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connects with the database using mysql or psql.",
	Run: func(cmd *cobra.Command, args []string) {
		databaseInstance := args[0]
		if len(args) < 1 {
			log.Fatal("Could not find value to add to databases!")
		}

		database.Connect(databaseInstance, databaseInput, passInput)
	},
}

func init() {
	connectCmd.Flags().StringVarP(&databaseInput, "database", "d", "", "Target Database to connect")
	connectCmd.MarkFlagRequired("database")

	connectCmd.Flags().StringVarP(&passwordInput, "password", "p", "", "Password to connect to the database")
}
