package database

import (
	"fmt"
	"log"
	"strings"

	"github.com/flow-cli/internal/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/viper"
)

func ListDatabases(areCredsHidden, isDecoded bool, databases []DatabaseConfig) {
	headers := getDatabaseConfigTableHeaders()
	var databaseRows []table.Row

	for _, db := range databases {
		db := handleDataVisibility(areCredsHidden, isDecoded, db)
		databaseRows = append(databaseRows, db.mapToTableRow())
	}

	utils.PrintTable(headers, databaseRows)
}

func FilterDatabases(query string, databases []DatabaseConfig) []DatabaseConfig {
	return databases
}

func isDbNameUnique(dbs []DatabaseConfig, name string) bool {
	for _, db := range dbs {
		if db.Name == name {
			return false
		}
	}

	return true
}

func AddDatabase(name string, host string, user string, password string, dbType string) {
	if name == "" {
		log.Fatal("Database cannot be an empty string!")
	}

	lowerCaseDbType := strings.ToLower(dbType)
	if lowerCaseDbType != "mysql" && lowerCaseDbType != "postgres" {
		log.Fatalf("\nCannot add database with type %s! Choose between 'mysql' or 'postgres'", dbType)
	}

	dbs := GetDatabasesFromConfig()
	isDbNameUnique := isDbNameUnique(dbs, name)
	if !isDbNameUnique {
		log.Fatalf("Failed to add database, '%s' is not unique!", name)
	}

	var database DatabaseConfig = DatabaseConfig{
		Name: name,
		Host: host,
		User: user,
		Pass: encodePassword(password),
		Type: lowerCaseDbType,
	}

	newDatabases := append(dbs, database)
	viper.Set("databases", newDatabases)
	viper.WriteConfig()

	database.Pass = HIDDEN_PASSWORD

	fmt.Println("Successfully added the following entry")
	utils.PrintTable(getDatabaseConfigTableHeaders(), []table.Row{database.mapToTableRow()})
}

func removeTargetDatabase(databases []DatabaseConfig, name string) []DatabaseConfig {
	var temp []DatabaseConfig

	for _, db := range databases {
		if db.Name != name {
			temp = append(temp, db)
		}
	}

	return temp
}

func DeleteDatabase(name string) {
	dbs := GetDatabasesFromConfig()

	if len(dbs) == 0 {
		log.Fatal("Can't delete database. No databases registered!")
	}

	updatedDbs := removeTargetDatabase(dbs, name)

	if len(updatedDbs) == len(dbs) {
		log.Fatalf("No database registered named: %s\n", name)
	}

	viper.Set("databases", updatedDbs)
	viper.WriteConfig()
	fmt.Printf("Successfully removed database: %s \n", name)
}
