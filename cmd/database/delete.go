/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package database

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func removeTargetDatabase(databases []DatabaseConfig, name string) []DatabaseConfig {
	var temp []DatabaseConfig

	for _, db := range databases {
		if db.Name != name {
			temp = append(temp, db)
		}
	}

	return temp
}

func deleteDatabase(name string) {
	dbs := getDatabasesFromConfig()

	if len(dbs) == 0 {
		log.Fatal("Can't delete database. No databases registered!")
	}

	updatedDbs := removeTargetDatabase(dbs, name)

	viper.Set("databases", updatedDbs)
	viper.WriteConfig()
	fmt.Printf("Successfully removed database: %s \n", name)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a database based on it's name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Could not find value to delete a database!")
		}

		deleteDatabase(args[0])
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
