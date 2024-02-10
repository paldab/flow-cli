/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package ip

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

type fullIP struct {
	Query       string `json:"query"`
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
	RegionName  string `json:"regionName"`
	City        string `json:"city"`
	ISP         string `json:"ISP"`
	TimeZone    string `json:"timeZone"`
}

var allFlag bool

func getIp() (fullIP, error) {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return fullIP{}, err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return fullIP{}, err
	}

	var ip fullIP
	json.Unmarshal(body, &ip)

	return ip, nil
}

func copyToClipboard(data string) {
	err := clipboard.Init()

	if err != nil {
		log.Fatal(err.Error())
	}

	clipboard.Write(clipboard.FmtText, []byte(data))
	clipboard.Read(clipboard.FmtImage)
}

func prettyPrint(data fullIP) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}

// ipCmd represents the ip command
var IpCmd = &cobra.Command{
	Use:   "ip",
	Short: "Gets current ip address",
	Long:  `Gets current ip address.`,
	Run: func(cmd *cobra.Command, args []string) {
		ip, err := getIp()
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		if allFlag {
			prettyPrint(ip)
			return
		}

		copyToClipboard(ip.Query)
		fmt.Println(ip.Query)
	},
}

func init() {
	IpCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "Get's more information about your IP")
}
