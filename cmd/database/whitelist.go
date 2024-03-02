package database

import (
	"log"
	"net"

	"github.com/flow-cli/internal/database"
	"github.com/spf13/cobra"
)

func isIpValid(ip string) bool {
	return net.ParseIP(ip) != nil
}

var networkName, project, instance string

var whitelistCmd = &cobra.Command{
	Use:   "whitelist <IP> -i <instance_name> -p <project>",
	Args:  cobra.ExactArgs(1),
	Short: "Whitelist the IP in the given database. Only works with Gcloud",
	Run: func(cmd *cobra.Command, args []string) {
		if !isIpValid(args[0]) {
			log.Fatal("Input is not a valid IP address")
		}
		database.Whitelist(args[0], project, instance, networkName)
	},
}

func init() {
	whitelistCmd.Flags().StringVarP(&instance, "instance", "i", "", "Name of the sql instance!")
	whitelistCmd.Flags().StringVarP(&project, "project", "p", "", "Google project of the sql instance!")
	whitelistCmd.Flags().StringVarP(&networkName, "name", "n", "", "Name of the whitelisted user!")

	whitelistCmd.MarkFlagRequired("instance")
	whitelistCmd.MarkFlagRequired("project")
	whitelistCmd.MarkFlagsRequiredTogether("instance", "project")
}
