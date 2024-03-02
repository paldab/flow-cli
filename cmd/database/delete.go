package database

import (
	"log"

	"github.com/flow-cli/internal/database"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a database based on it's name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Could not find value to delete a database!")
		}

		database.DeleteDatabase(args[0])
	},
}
