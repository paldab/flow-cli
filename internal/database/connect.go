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
)

func createDatabaseConnectCommand(dbType, host, user, password, database string) string {
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

func handlePasswordDecryption(password string) string {
	if isBase64Encoded(password) {
		return decodePassword(password)
	}

	return password
}

func Connect(databaseInstance, targetDatabase, passwordInput string) {
	if targetDatabase == "" {
		log.Fatal("Must specify a database")
	}

	targetDb, err := dbLookup(databaseInstance)
	if err != nil {
		log.Fatal(err.Error())
	}

	if passwordInput != "" {
		targetDb.Pass = passwordInput
	}

	decryptedPassword := handlePasswordDecryption(targetDb.Pass)

	query := createDatabaseConnectCommand(targetDb.Type, targetDb.Host, targetDb.User, decryptedPassword, targetDatabase)
	args := strings.Split(query, " ")
	connectToDatabaseCli(args)
}
