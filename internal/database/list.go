package database

import (
	"cmp"
	"log"
	"slices"
	"strings"

	"github.com/flow-cli/internal/utils"
	"github.com/jedib0t/go-pretty/v6/table"
)

func ListDatabases(query string, areCredsHidden, isDecoded bool, databases []DatabaseConfig) {
	var databaseRows []table.Row

	filteredDatabases := FilterDatabases(query, databases)

	sortByName := func(a, b DatabaseConfig) int {
		return cmp.Compare(a.Name, b.Name)
	}

	slices.SortFunc(filteredDatabases, sortByName)
	for _, db := range filteredDatabases {
		db := handleDataVisibility(isDecoded, db)
		databaseRows = append(databaseRows, db.mapToTableRow(areCredsHidden))
	}

	headers := getDatabaseConfigTableHeaders(areCredsHidden)
	utils.PrintTable(headers, databaseRows)
}

func FilterDatabases(query string, databases []DatabaseConfig) []DatabaseConfig {
	if query == "" {
		return databases
	}

	var result []DatabaseConfig
	for _, db := range databases {
		if strings.Contains(db.Name, query) {
			result = append(result, db)
		}
	}

	if len(result) == 0 {
		log.Fatal("Could not find any databases with this name!")
	}

	return result
}
