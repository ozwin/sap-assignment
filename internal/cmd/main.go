package main

import (
	"flag"
	"fmt"
	"os"

	config "github.com/ozwin/sap-assignment/internal/app/configs"
	"github.com/ozwin/sap-assignment/internal/app/models"
	service "github.com/ozwin/sap-assignment/internal/app/services"
)

func main() {
	hasBikeTrail := flag.Bool("has_bike_trail", false, "Filter trails that have a bike trail")
	hasCamping := flag.Bool("has_camping", false, "Filter trails that have camping facilities")
	hasHikingTrail := flag.Bool("has_hiking_trail", false, "Filter trails that have hiking trails")
	hasDogCompost := flag.Bool("has_dog_compost", false, "Filter trails that have dog compost bins")
	hasFishing := flag.Bool("has_fishing", false, "Filter trails that have fishing facilities")
	hasPicnic := flag.Bool("has_picnic", false, "Filter trails that have picnic areas")
	hasFees := flag.Bool("has_fees", false, "Filter trails that have fees")
	hasRecycleBin := flag.Bool("has_recycle_bin", false, "Filter trails that have recycle bins")
	flag.Parse()
	//intialize application
	trails, err := models.NewTrailsStore(config.TrailsFileName)
	if err != nil {

		fmt.Fprintf(os.Stderr, "Failed to load the CSV file, Please check the path: %v\n", err)
		os.Exit(1)
	}

	trailService := service.NewTrailService(trails)
	filters := make(map[service.FilterField]interface{})

	if *hasBikeTrail {
		filters[service.HasBikeTrail] = true
	}
	if *hasCamping {
		filters[service.HasCamping] = true
	}
	if *hasHikingTrail {
		filters[service.HasHikingTrail] = true
	}
	if *hasDogCompost {
		filters[service.HasDogCompost] = true
	}
	if *hasFishing {
		filters[service.HasFishing] = true
	}
	if *hasPicnic {
		filters[service.HasPicnic] = true
	}
	if *hasFees {
		filters[service.HasFees] = true
	}
	if *hasRecycleBin {
		filters[service.HasRecycleBin] = true
	}

	filteredTrails, err := trailService.FilterTrails(filters)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error filtering trails: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Filtered Trails:")
	for _, trail := range filteredTrails {
		fmt.Printf("- %s (%s) - Difficulty: %s\n", trail.AccessName, trail.Address, trail.Difficulty)
	}
}
