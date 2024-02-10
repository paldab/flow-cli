package database

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getDatabasesFromConfig() []DatabaseConfig {
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err.Error())
	}

	var dbs []DatabaseConfig
	if err := viper.UnmarshalKey("databases", &dbs); err != nil {
		log.Fatal(err.Error())
	}

	return dbs
}

func dbLookup(target string) DatabaseConfig {
	dbs := getDatabasesFromConfig()
	for _, db := range dbs {
		if db.Name == target {
			return db
		}
	}

	log.Fatal("Could not find database! Make sure you registered the database")
	return DatabaseConfig{}
}

var DatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "database represents the database module of the CLI. Managing your databases for easy overview and connection",
	Long:  `database represents the database module of the CLI. Managing your databases for easy overview and connection`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	DatabaseCmd.Aliases = []string{"db"}

	DatabaseCmd.AddCommand(addCmd)
	DatabaseCmd.AddCommand(listCmd)
	DatabaseCmd.AddCommand(deleteCmd)
	DatabaseCmd.AddCommand(whitelistCmd)
	DatabaseCmd.AddCommand(connectCmd)
}
