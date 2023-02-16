package util

import "fmt"

// WhoisApiSender api sender for get where user came from
func WhoisApiSender() {

	// Whois api key
	apiKey := CheckAndGetApiKey("WHOIS_API_KEY")

	fmt.Println(apiKey)
}
