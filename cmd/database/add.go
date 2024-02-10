/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package database

import (
	"flow/cli/utils"
	"fmt"
	"log"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var host string
var user string
var pass string
var dbType string

func isDbNameUnique(dbs []DatabaseConfig, name string) bool {
	for _, db := range dbs {
		if db.Name == name {
			return false
		}
	}

	return true
}

// TODO: Encrypt password
func addDatabase(name string, host string, user string, password string, dbType string) {
	if name == "" {
		log.Fatal("Database cannot be an empty string!")
	}

	lowerCaseDbType := strings.ToLower(dbType)
	if lowerCaseDbType != "mysql" && lowerCaseDbType != "postgres" {
		log.Fatal(fmt.Sprintf("\nCannot add database with type %s! Choose between 'mysql' or 'postgres'", dbType))
	}

	dbs := getDatabasesFromConfig()
	isDbNameUnique := isDbNameUnique(dbs, name)
	if !isDbNameUnique {
		log.Fatal(fmt.Sprintf("Failed to add database, '%s' is not unique!", name))
	}

	var database DatabaseConfig = DatabaseConfig{
		Name: name,
		Host: host,
		User: user,
		Pass: pass,
		Type: lowerCaseDbType,
	}

	newDatabases := append(dbs, database)
	viper.Set("databases", newDatabases)
	viper.WriteConfig()

	database.Pass = "****"

	fmt.Println("Successfully added the following entry")
	utils.PrintTable(getDatabaseConfigTableHeaders(), []table.Row{database.mapToTableRow()})
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a database to the config",
	Long: `Example usage 

database add example-prod --host=0.0.0.0 -u user1 -p pass1`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Could not find value to add to databases!")
		}

		addDatabase(args[0], host, user, pass, dbType)
	},
}

func init() {
	addCmd.Flags().StringVar(&host, "host", "", "Define the host of the database")
	addCmd.MarkFlagRequired("host")

	addCmd.Flags().StringVarP(&user, "user", "u", "", "Define a user for the database")
	addCmd.MarkFlagRequired("user")

	addCmd.Flags().StringVarP(&pass, "password", "p", "", "Define a password for the database")

	addCmd.Flags().StringVarP(&dbType, "type", "t", "", "Type of database. MySQL or Postgres")
	addCmd.MarkFlagRequired("type")
}
