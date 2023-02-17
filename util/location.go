package util

import (
	"encoding/csv"
	"os"
	"strings"
)

var locations = make(map[string]string)

func LocationCode(addr string) string {
	if addr == "" {
		return "no addr"
	}
	city := strings.Split(addr, " ")
	for index, word := range city {
		if index == 1 {
			if locations[word[:len(word)-3]] != "" {
				return locations[word[:len(word)-3]]
			}
		}
		if locations[word] != "" {
			return locations[word]
		}
	}
	return "no addr"
}

// SetLocationCodeMap set map for location codes
func SetLocationCodeMap() {
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
