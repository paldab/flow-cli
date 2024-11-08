package ip

import (
	"fmt"
	"log"

	"github.com/flow-cli/internal/network"
	"github.com/spf13/cobra"
)

var allFlag bool

var IpCmd = &cobra.Command{
	Use:   "ip",
	Short: "Gets current ip address",
	Run: func(cmd *cobra.Command, args []string) {
		ip, err := network.GetIp()
		if err != nil {
			log.Fatalf("could not get IP. %s", err.Error())
		}

		if allFlag {
            if err := network.PrettyPrint(ip); err != nil {
                log.Fatal(err.Error())
            }
			return
		}

		fmt.Println(ip.Query)
	},
}

func init() {
	IpCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "Get's more information about your IP")
}
