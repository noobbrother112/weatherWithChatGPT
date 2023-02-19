package main

import (
	"fmt"
	"github.com/noobbrother112/weatherWithChatGPT/util"
)

func main() {
	//util.GptApiSender()
	util.SetLocationCodeMap()
	addr := util.WhoisApiSender("220.72.89.208")
	fmt.Println(addr)
	code := util.LocationCode(addr)
	fmt.Println(code)
}
