package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var locations = make(map[string]string)

func main() {
	//util.GptApiSender()
	//util.WhoisApiSender({your ip})
	setLocationCodeMap()
	fmt.Println(locations["백령도"])
}

// set map for location codes
func setLocationCodeMap() {
	file, err := os.Open("./data/locationCode.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		locations[record[0]] = record[1]
	}
}
