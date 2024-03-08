package database

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

const HIDDEN_PASSWORD = "********"

type DatabaseConfig struct {
	Name string
	Host string
	User string
	Pass string
	Type string
}

func getDatabaseConfigTableHeaders(areCredsHidden bool) table.Row {
	if areCredsHidden {
		return table.Row{"Name", "Host", "Type"}
	}

	return table.Row{"Name", "Host", "User", "Password", "Type"}
}

func (d DatabaseConfig) mapToTableRow(areCredsHidden bool) table.Row {
	if areCredsHidden {
		return table.Row{d.Name, d.Host, d.Type}
	}

	return table.Row{d.Name, d.Host, d.User, d.Pass, d.Type}
}
