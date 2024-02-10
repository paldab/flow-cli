package database

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

type DatabaseConfig struct {
	Name string
	Host string
	User string
	Pass string
	Type string
}

func getDatabaseConfigTableHeaders() table.Row {
	return table.Row{"Name", "Host", "User", "Password", "Type"}
}

func (d DatabaseConfig) mapToTableRow() table.Row {
	return table.Row{d.Name, d.Host, d.User, d.Pass, d.Type}
}
