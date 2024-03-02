package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/flow-cli/cmd/database"
	"github.com/flow-cli/cmd/deploy"
	"github.com/flow-cli/cmd/generate"
	"github.com/flow-cli/cmd/ip"
	"github.com/flow-cli/cmd/kubernetes"
	internalDatabase "github.com/flow-cli/internal/database"
	"github.com/flow-cli/internal/task"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var verbose bool

const DEFAULTCONFIGFOLDER = ".flow"
const DEFAULTCONFIGFILE = "config.yaml"

var rootCmd = &cobra.Command{
	Use:   "flow",
	Short: "Flow is a CLI tool to automate repetitive tasks",
	Long:  `Flow is a CLI tool to automate repetitive tasks`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func handleDefaultConfigSetup(folderPath string, fileName string) {
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

func handleDefaultCreateDatabaseConfig() {
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

func handleDefaultCreateTaskHistory(folderPath string) {
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

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(database.DatabaseCmd)
	rootCmd.AddCommand(ip.IpCmd)
	rootCmd.AddCommand(deploy.DeployCmd)
	rootCmd.AddCommand(kubernetes.KubernetesCmd)
	rootCmd.AddCommand(generate.GenerateCmd)
	// rootCmd.AddCommand(task.TaskCmd)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.flow/config.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logs")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		configFolderLocation := path.Join(home, DEFAULTCONFIGFOLDER)
		handleDefaultConfigSetup(configFolderLocation, DEFAULTCONFIGFILE)
		handleDefaultCreateDatabaseConfig()
		// handleDefaultCreateTaskHistory(configFolderLocation)

		viper.WriteConfig()
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if verbose {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	}

	if !verbose {
		log.SetFlags(0)
	}
}
