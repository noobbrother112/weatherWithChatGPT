package util

import (
	"encoding/json"
	"fmt"
)

const whoisEndpoint = "http://apis.data.go.kr/B551505/whois/ip_address?"

// WhoisApiSender api sender for get where user came from
func WhoisApiSender(ip string) string {
	c := make(chan []byte)

	// Whois api key
	apiKey := CheckAndGetApiKey("WHOIS_API_KEY")
	whoisUrl := whoisEndpoint + "serviceKey=" + apiKey + "&query=" + ip + "&answer=json"

	go ApiSender("GET", whoisUrl, apiKey, nil, c)
	var body = <-c

	// convert json to map
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	responsMap := data["response"].(map[string]interface{})
	whois := responsMap["whois"].(map[string]interface{})
	korean := whois["korean"].(map[string]interface{})
	if korean["user"] != nil {
		user := korean["user"].(map[string]interface{})
		netinfo := user["netinfo"].(map[string]interface{})
		addr := netinfo["addr"].(string)

		return addr
	} else {
		// 유저 정보가 조회되지 않음
		return ""
	}

}
