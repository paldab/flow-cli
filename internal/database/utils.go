package database

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

func GetDatabasesFromConfig() []DatabaseConfig {
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err.Error())
	}

	var dbs []DatabaseConfig
	if err := viper.UnmarshalKey("databases", &dbs); err != nil {
		log.Fatal(err.Error())
	}

	return dbs
}

func dbLookup(target string) (DatabaseConfig, error) {
	dbs := GetDatabasesFromConfig()
	for _, db := range dbs {
		if db.Name == target {
			return db, nil
		}
	}

	return DatabaseConfig{}, errors.New("could not find database! Make sure that you registered the database")
}
