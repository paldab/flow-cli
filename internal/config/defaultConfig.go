package config

import (
	"log"
	"os"
	"path"

	internalDatabase "github.com/flow-cli/internal/database"
	"github.com/flow-cli/internal/task"
	"github.com/spf13/viper"
)

func HandleDefaultConfigSetup(folderPath string, fileName string) {
	_, err := os.Stat(folderPath)

	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err.Error())
	}

	if os.IsNotExist(err) {
		os.Mkdir(folderPath, os.ModePerm)
		os.Create(path.Join(folderPath, fileName))
	}

	configFileLocation := path.Join(folderPath, fileName)
	_, err = os.Stat(configFileLocation)
	if err != nil {
		if os.IsNotExist(err) {
			os.Create(configFileLocation)
		} else {
			log.Fatal(err.Error())
		}
	}

	viper.AddConfigPath(folderPath)
	viper.SetConfigType("yaml")
	viper.SetConfigName(fileName)
}

func HandleDefaultCreateDatabaseConfig() {
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err.Error())
	}

	var dbs []internalDatabase.DatabaseConfig
	if err := viper.UnmarshalKey("databases", &dbs); err != nil {
		log.Fatal(err.Error())
	}

	if len(dbs) == 0 {
		viper.SetDefault("databases", []internalDatabase.DatabaseConfig{})
	}
}

func HandleDefaultCreateTaskHistory(folderPath string) {
	fileName := "flowtaskhistory.csv"
	targetFilePath := path.Join(folderPath, fileName)
	_, err := os.Stat(targetFilePath)

	if err != nil {
		if os.IsNotExist(err) {
			os.Create(targetFilePath)
		} else {
			log.Fatal(err.Error())
		}
	}

	viper.SetDefault("tasks.history.path", targetFilePath)
	task.SetTaskHeaders()
}
