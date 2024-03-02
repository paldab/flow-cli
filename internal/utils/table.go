package utils

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func PrintTable(header table.Row, rows []table.Row) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(header)
	t.AppendRows(rows)
	t.Render()
}
