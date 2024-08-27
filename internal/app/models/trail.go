package models

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Trail struct {
	AccessName string `json:"access_name,omitempty"`
	Address    string `json:"address,omitempty"`
	Difficulty string `json:"difficulty,omitempty"`
	Fees       bool   `json:"fees,omitempty"`
	HorseTrail string `json:"horse_trail,omitempty"`
	Picnic     *bool  `json:"picnic,omitempty"`
}

type Trails []Trail

func readTrailsDataFromCSV(fileName string) (Trails, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	var trails Trails
	headerMap := make(map[string]int)
	for index, col := range rows[0] {
		headerMap[col] = index
	}

	for _, row := range rows[1:] {
		trail := Trail{
			AccessName: strings.TrimSpace(row[headerMap["AccessName"]]),
			Address:    strings.TrimSpace(row[headerMap["Address"]]),
			Difficulty: strings.TrimSpace(row[headerMap["ADAtrail"]])}
		trails = append(trails, trail)
	}

	return trails, nil
}
