package network

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func GetIp() (fullIP, error) {
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

func PrettyPrint(data fullIP) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}