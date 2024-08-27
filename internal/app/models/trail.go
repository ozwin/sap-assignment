package models

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/ozwin/sap-assignment/internal/utils"
)

type Trail struct {
	// for display
	AccessName string `json:"access_name,omitempty"`
	Address    string `json:"address,omitempty"`
	Difficulty string `json:"difficulty,omitempty"`

	//for filtering
	HasBikeTrail   bool `json:"has_bike_trail,omitempty"`
	HasCamping     bool `json:"has_camping,omitempty"`
	HasHikingTrail bool `json:"has_hiking,omitempty"`
	// HorseTrail   bool `json:"has_horse_trail,omitempty"`
	//for options

	HasFees       bool `json:"has_fees,omitempty"`
	HasFishing    bool `json:"has_fishing,omitempty"`
	HasPicnic     bool `json:"has_picnic,omitempty"`
	HasRecycleBin bool `json:"has_recycle_bin,omitempty"`
	HasDogCompost bool `json:"has_dog_compost,omitempty"`
}

type Trails []Trail

func NewTrailsStore(fileName string) (Trails, error) {
	return readTrailsDataFromCSV(fileName)
}

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
			AccessName:     strings.TrimSpace(row[headerMap["AccessName"]]),
			Address:        strings.TrimSpace(row[headerMap["Address"]]),
			Difficulty:     strings.TrimSpace(row[headerMap["ADAtrail"]]),
			HasFees:        utils.StringToBooleanMapper(row[headerMap[" "]]),
			HasFishing:     utils.StringToBooleanMapper(row[headerMap["FISHING"]]),
			HasPicnic:      utils.StringToBooleanMapper(row[headerMap["PICNIC"]]),
			HasRecycleBin:  utils.StringToBooleanMapper(row[headerMap["RecycleBin"]]),
			HasDogCompost:  utils.StringToBooleanMapper(row[headerMap["DogCompost"]]),
			HasBikeTrail:   utils.StringToBooleanMapper(row[headerMap["BikeTrail"]]),
			HasCamping:     utils.StringToBooleanMapper(row[headerMap["ADAcamping"]]),
			HasHikingTrail: utils.CheckIfTrailExists(row[headerMap["ADAtrail"]]),
		}
		trails = append(trails, trail)
	}

	return trails, nil
}
