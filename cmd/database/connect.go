/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package database

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var passwordInput string
var database string

func CreateDatabaseConnectCommand(dbType, host, user, password, database string) string {
	switch dbType {
	case "mysql":
		return fmt.Sprintf("mysql --host=%s --user=%s --password=%s %s", host, user, password, database)
	case "postgres":
		return fmt.Sprintf("psql postgresql://%s:%s@%s/%s", user, password, host, database)
	default:
		return ""
	}
}

func connectToDatabaseCli(args []string) {
	var attr os.ProcAttr
	attr.Files = []*os.File{
		os.Stdin,
		os.Stdout,
		os.Stderr,
	}

	path, err := exec.LookPath(args[0])
	if err != nil {
		log.Fatal(err.Error())
	}

	process, err := os.StartProcess(path, args, &attr)
	if err != nil {
		log.Fatal(err.Error())
	}

	process.Wait()
}

func connect(db string) {
	targetDb := dbLookup(db)

	if database == "" {
		log.Fatal("Must specify a database")
	}

	if passwordInput != "" {
		targetDb.Pass = passwordInput
	}

	query := CreateDatabaseConnectCommand(targetDb.Type, targetDb.Host, targetDb.User, targetDb.Pass, database)
	args := strings.Split(query, " ")
	connectToDatabaseCli(args)
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connects with the database using mysql or psql.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Could not find value to add to databases!")
		}

		connect(args[0])
	},
}

func init() {
	connectCmd.Flags().StringVarP(&database, "database", "d", "", "Target Database to connect")
	connectCmd.MarkFlagRequired("database")

	connectCmd.Flags().StringVarP(&passwordInput, "password", "p", "", "Password to connect to the database")
}
