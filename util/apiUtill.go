package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// ApiSender sending api and get json back
func ApiSender(method, url, apiKey string, jsonStr []byte, c chan<- []byte) {
	fmt.Println(string(jsonStr))

	// Create HTTP request
	req, err := http.NewRequest(method, url, strings.NewReader(string(jsonStr)))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	c <- body
}

func CheckAndGetApiKey(key string) string {
	apiKey := os.Getenv(key)
	if apiKey == "" {
		log.Fatalln("WHOIS API KEY is not set")
	}
	return apiKey
}
