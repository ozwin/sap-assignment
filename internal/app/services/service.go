package service

import (
	"github.com/ozwin/sap-assignment/internal/app/models"
)

type TrailService struct {
	trails models.Trails
}

type FilterField string

const (
	HasBikeTrail   FilterField = "HasBikeTrail"
	HasCamping     FilterField = "HasCamping"
	HasHikingTrail FilterField = "HasHikingTrail"
	HasDogCompost  FilterField = "HasDogCompost"
	HasFishing     FilterField = "HasFishing"
	HasPicnic      FilterField = "HasPicnic"
	HasFees        FilterField = "HasFees"
	HasRecycleBin  FilterField = "HasRecycleBin"
)

func NewTrailService(trails models.Trails) *TrailService {
	// trails, err := models.NewTrailsStore()
	// if err != nil {
	// 	panic("Failed to get the trails, shutting down the application")
	// }
	return &TrailService{
		trails: trails,
	}
}

func (ts *TrailService) GetAll() models.Trails {
	return ts.trails
}

func (ts *TrailService) FilterTrails(filters map[FilterField]interface{}) (models.Trails, error) {
	if len(filters) == 0 {
		return ts.trails, nil
	}
	var result models.Trails
	for _, trail := range ts.trails {
		match := true

		for filter, value := range filters {
			switch filter {
			case HasBikeTrail:
				if val, ok := value.(bool); ok && trail.HasBikeTrail != val {
					match = false
				}
			case HasCamping:
				if val, ok := value.(bool); ok && trail.HasCamping != val {
					match = false
				}
			case HasHikingTrail:
				if val, ok := value.(bool); ok && trail.HasHikingTrail != val {
					match = false
				}
			case HasDogCompost:
				if val, ok := value.(bool); ok && trail.HasDogCompost != val {
					match = false
				}
			case HasFishing:
				if val, ok := value.(bool); ok && trail.HasFishing != val {
					match = false
				}
			case HasPicnic:
				if val, ok := value.(bool); ok && trail.HasPicnic != val {
					match = false
				}
			case HasFees:
				if val, ok := value.(bool); ok && trail.HasFees != val {
					match = false
				}
			case HasRecycleBin:
				if val, ok := value.(bool); ok && trail.HasRecycleBin != val {
					match = false
				}
			}
			if !match {
				break
			}
		}
		if match {
			result = append(result, trail)
		}
	}
	return result, nil
}
