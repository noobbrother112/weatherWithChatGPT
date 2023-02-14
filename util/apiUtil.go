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
	// OpenAI API key
	apiKey := checkApiKey()

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

	// Create HTTP request
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(string(reqJson)))
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

func checkApiKey() string {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("API key is not set")
	}
	return apiKey
}
