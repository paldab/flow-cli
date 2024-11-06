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
	config "github.com/flow-cli/internal/config"
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
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
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
		config.HandleDefaultConfigSetup(configFolderLocation, DEFAULTCONFIGFILE)
		config.HandleDefaultCreateDatabaseConfig()
		// config.HandleDefaultCreateTaskHistory(configFolderLocation)

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
