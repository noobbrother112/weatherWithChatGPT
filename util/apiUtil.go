package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const endpoint = "https://api.openai.com/v1/completions"

type GptRequest struct {
	Prompt           string  `json:"prompt"`
	Model            string  `json:"model"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      float32 `json:"temperature"`
	FrequencyPenalty float32 `json:"frequency_penalty"`
	PresencePenalty  float32 `json:"presence_penalty"`
}

type GptResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

// GptApiSender for ChatGPT
func GptApiSender() {
	//Channel for http response
	c := make(chan []byte)

	// OpenAI API key
	apiKey := checkGPTApiKey()

	// Set up request
	request := GptRequest{
		Model:            "text-davinci-003",
		Prompt:           "영하10도에선 어떤 옷을 입어야 하니?",
		MaxTokens:        600,
		Temperature:      0,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}

	// Convert request to JSON
	reqJson, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(reqJson))

	go apiSender("POST", endpoint, apiKey, reqJson, c)

	var body = <-c
	fmt.Println(string(body))
	// Convert response to JSON
	var response GptResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print response
	fmt.Println(response.Choices[0].Text)
}

func checkGPTApiKey() string {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("API key is not set")
	}
	return apiKey
}

func apiSender(method, url, apiKey string, jsonStr []byte, c chan<- []byte) {
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
