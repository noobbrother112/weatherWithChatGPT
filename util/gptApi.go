package util

import (
	"encoding/json"
	"fmt"
)

const gptEndpoint = "https://api.openai.com/v1/completions"

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
	apiKey := CheckAndGetApiKey("API_KEY")

	// Set up request
	request := GptRequest{
		Model:            "text-davinci-003",
		Prompt:           "이번예보기간에는고기압의가장자리에들어가끔구름많겠습니다. 기온은평년(최저기온 : 2 ~ 7도, 최고기온 : 16 ~ 19도)과비슷하거나조금높겠습니다. 강수량은평년(강수량 : 1~4mm)보다적겠습니다. 서해중부해상의물결은 1~2m로일겠습니다. 이런날씨에는 어떤 옷을 입어야하는지 상의, 하의, 신발 순으로 추천해주고 만약 준비물이 필요하다면 그것도 추천해줘",
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

	go ApiSender("POST", gptEndpoint, apiKey, reqJson, c)
	var body = <-c

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
