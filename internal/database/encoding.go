package database

import (
	"encoding/base64"
	"log"
)

func isBase64Encoded(input string) bool {
	_, err := base64.StdEncoding.DecodeString(input)
	return err == nil
}

func encodePassword(data string) string {
	secret := base64.StdEncoding.EncodeToString([]byte(data))
	return secret
}

func decodePassword(data string) string {
	decode, err := base64.StdEncoding.DecodeString(data)

	if err != nil {
		log.Fatal("Could not decode password!")
	}

	return string(decode)
}
