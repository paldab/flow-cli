/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package task

import (
	"io"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

// createBackupCmd represents the createBackup command
var createBackupCmd = &cobra.Command{
	Use:   "createBackup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		sourceFile := getTaskHistoryFile()
		parentDir := path.Dir(sourceFile.Name())
		backupPrefix := ".backup"
		backupFileName := strings.Split(TASKHISTORYFILE, ".")[0] + backupPrefix + ".csv"

		backupFile, err := os.Create(path.Join(parentDir, backupFileName))

		if err != nil {
			log.Fatal(err.Error())
		}

		_, err = io.Copy(backupFile, sourceFile)

		if err != nil {
			log.Fatal(err.Error())
		}

		println("Backup of tasks has been created!")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createBackupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createBackupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
