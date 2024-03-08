/*
This module is separate from the registered databases. This is to whitelist a ip to a sql instance
*/
package database

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type authorizedNetwork struct {
	Kind  string `json:"kind"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TargetJSON struct {
	Settings struct {
		IPConfiguration struct {
			AuthorizedNetworks []authorizedNetwork `json:"authorizedNetworks"`
		} `json:"ipConfiguration"`
	} `json:"settings"`
}

func getLocalGcloudPath() string {
	val, err := exec.LookPath("gcloud")
	if err != nil {
		log.Fatal(err.Error())
	}

	return val
}

func getAuthToken(gcloud string) string {
	cmd := exec.Command(gcloud, "auth", "application-default", "print-access-token")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing gcloud command: %v\n", err)
	}

	token := strings.TrimSpace(string(output))
	if token == "" {
		log.Fatal("Access token is empty!")
	}

	return token
}

func makeAPIGetRequest(token, url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	bearer := "Bearer " + token
	req.Header.Set("Authorization", bearer)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}

	return res
}

func handleNetworksFromJson(out string) []authorizedNetwork {
	var jsonData map[string]interface{}
	if err := json.Unmarshal([]byte(out), &jsonData); err != nil {
		log.Fatal(err.Error())
	}

	settings, exists := jsonData["settings"].(map[string]interface{})
	if !exists {
		log.Fatal("Settings key not found during unmarshaling authnetworks")
	}

	config, exists := settings["ipConfiguration"].(map[string]interface{})
	if !exists {
		log.Fatal("ipConfiguration key not found during unmarshaling authnetworks")
	}

	authNetworksMap, exists := config["authorizedNetworks"].([]interface{})
	if !exists {
		log.Fatal("authorizedNetworks key not found or not an array")
	}

	var networks []authorizedNetwork
	for _, network := range authNetworksMap {
		if networkMap, ok := network.(map[string]interface{}); ok {
			networks = append(networks, authorizedNetwork{
				Name:  networkMap["name"].(string),
				Value: networkMap["value"].(string),
			})
		} else {
			log.Fatal("Could not convert existing networks!")
		}
	}

	return networks
}

func getExistingNetworks(token, project, instance string) []authorizedNetwork {
	url := fmt.Sprintf("https://sqladmin.googleapis.com/v1/projects/%s/instances/%s", project, instance)
	res := makeAPIGetRequest(token, url)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	result := string([]byte(body))
	return handleNetworksFromJson(result)
}

func addNewNetwork(token string, networks []authorizedNetwork, project, instance string) {
	url := fmt.Sprintf("https://sqladmin.googleapis.com/v1/projects/%s/instances/%s", project, instance)

	var networkJson TargetJSON
	networkJson.Settings.IPConfiguration.AuthorizedNetworks = networks

	jsonPayload, err := json.Marshal(networkJson)
	if err != nil {
		log.Fatal(err.Error())
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatal(err.Error())
	}

	bearer := "Bearer " + token
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Body.Close()

	fmt.Printf("%v\n", res.Status)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}
	fmt.Printf("Response Body:\n%s\n", body)

	if res.StatusCode != 200 {
		log.Fatal("Something went wrong when adding the whitelist!")
	}
}

func Whitelist(ip, project, instance, networkName string) {
	gcloud := getLocalGcloudPath()
	token := getAuthToken(gcloud)
	networks := getExistingNetworks(token, project, instance)
	networks = append(networks, authorizedNetwork{
		Name:  networkName,
		Value: ip,
	})

	addNewNetwork(token, networks, project, instance)
}
