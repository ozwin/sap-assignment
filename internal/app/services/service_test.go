package service

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/ozwin/sap-assignment/internal/app/models"
)

func TestTrailService_GetAll(t *testing.T) {
	tests := []struct {
		name string
		ts   *TrailService
		want models.Trails
	}{
		// TODO: Add test cases.
		//Not part of the scope
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ts.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrailService.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrailService_FilterTrails(t *testing.T) {
	//creating sample data for tests
	fileName := "sample_data"
	createSampledataFile(t, fileName)
	defer deleteSampleDataFile(t, fileName)

	//initalizing service instance
	trails, _ := models.NewTrailsStore(fileName)
	ts := NewTrailService(trails)

	//filters
	//case 1: No filters
	//case 2: only bike trails case
	bikeTrailFilter := make(map[FilterField]interface{})
	bikeTrailFilter[HasBikeTrail] = true

	//case 2: multiple filter scenarios
	mutlipleFilters := make(map[FilterField]interface{})
	mutlipleFilters[HasHikingTrail] = true
	mutlipleFilters[HasPicnic] = true

	type args struct {
		filters map[FilterField]interface{}
	}
	tests := []struct {
		name    string
		ts      *TrailService
		args    args
		want    models.Trails
		wantErr bool
	}{
		// TODO: Add test cases.
		{

			name: "Valid Data with no filters",
			ts:   ts,
			args: args{},
			want: models.Trails{{
				Address:        "621 Flagstaff Summit Rd",
				AccessName:     "Flagstaff Summit West",
				Difficulty:     "Moderate",
				HasHikingTrail: true,
				HasBikeTrail:   false,
				HasCamping:     false,
				HasFees:        false,
				HasFishing:     false,
				HasPicnic:      true,
				HasRecycleBin:  true,
				HasDogCompost:  false,
			}, {
				Address:        "790 Flagstaff Summit Rd",
				AccessName:     "Flagstaff Summit East",
				Difficulty:     "Difficult",
				HasHikingTrail: true,
				HasBikeTrail:   false,
				HasCamping:     false,
				HasFees:        false,
				HasFishing:     false,
				HasPicnic:      true,
				HasRecycleBin:  true,
				HasDogCompost:  false,
			}, {
				Address:        "4705 95th St",
				AccessName:     "East Boulder Trail at White Rocks",
				Difficulty:     "No",
				HasHikingTrail: false,
				HasBikeTrail:   true,
				HasCamping:     false,
				HasFees:        false,
				HasFishing:     false,
				HasPicnic:      false,
				HasRecycleBin:  false,
				HasDogCompost:  false,
			},
			},
			wantErr: false,
		},
		{

			name: "Only Bike trails",
			ts:   ts,
			args: args{filters: bikeTrailFilter},
			want: models.Trails{{
				Address:        "4705 95th St",
				AccessName:     "East Boulder Trail at White Rocks",
				Difficulty:     "No",
				HasHikingTrail: false,
				HasBikeTrail:   true,
				HasCamping:     false,
				HasFees:        false,
				HasFishing:     false,
				HasPicnic:      false,
				HasRecycleBin:  false,
				HasDogCompost:  false,
			},
			},
			wantErr: false,
		}, {

			name: "Multiple filter with mulitple options ",
			ts:   ts,
			args: args{filters: mutlipleFilters},
			want: models.Trails{{
				Address:        "621 Flagstaff Summit Rd",
				AccessName:     "Flagstaff Summit West",
				Difficulty:     "Moderate",
				HasHikingTrail: true,
				HasBikeTrail:   false,
				HasCamping:     false,
				HasFees:        false,
				HasFishing:     false,
				HasPicnic:      true,
				HasRecycleBin:  true,
				HasDogCompost:  false,
			}, {
				Address:        "790 Flagstaff Summit Rd",
				AccessName:     "Flagstaff Summit East",
				Difficulty:     "Difficult",
				HasHikingTrail: true,
				HasBikeTrail:   false,
				HasCamping:     false,
				HasFees:        false,
				HasFishing:     false,
				HasPicnic:      true,
				HasRecycleBin:  true,
				HasDogCompost:  false,
			},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ts.FilterTrails(tt.args.filters)
			if (err != nil) != tt.wantErr {
				t.Errorf("TrailService.FilterTrails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrailService.FilterTrails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createSampledataFile(t *testing.T, fileName string) {
	//ToDo: check later for this file existence to avoid redundency
	testFile, err := os.Create(fileName)
	if err != nil {
		t.Fatalf("error creating temp file for testing %v", err)
	}
	defer testFile.Close()
	data := `FID,RESTROOMS,PICNIC,FISHING,AKA,AccessType,AccessID,Class,Address,Fee,BikeRack,BikeTrail,DogTube,Grills,TrashCans,ParkSpaces,ADAsurface,ADAtoilet,ADAfishing,ADAcamping,ADApicnic,ADAtrail,ADAparking,ADAfacilit,ADAfacName,HorseTrail,DateFrom,DateTo,RecycleBin,DogCompost,AccessName,THLeash
			0,Yes,Yes,No, ,TH,279,T3,621 Flagstaff Summit Rd,Yes,No,No,1,Yes,4,12,Asphalt,Yes,No,No,Yes,Moderate,Yes,Yes,Wood Shelter,Not Recommended,12/31/2005 0:00,12/31/2099 0:00,Yes,No,Flagstaff Summit West,Yes
			1,Yes,Yes,No, ,TH,277,T3,790 Flagstaff Summit Rd,Yes,Yes,No,1,Yes,6,56,Asphalt,Yes,No,No,Yes,Difficult,Yes,Yes,Nature Center,Not Recommended,12/31/2005 0:00,12/31/2099 0:00,Yes,No,Flagstaff Summit East,Yes
			2,No,No,No, ,TH,502a,T1,4705 95th St,No,Yes,Yes,0,No,1,6,No,No,No,No,No,No,No,No, ,Not Recommended,12/31/2005 0:00,12/31/2099 0:00,No,No,East Boulder Trail at White Rocks,Yes`
	if _, err := testFile.Write([]byte(data)); err != nil {
		t.Fatalf("error writing to temp file: %v", err)
	}
}

func deleteSampleDataFile(t *testing.T, fileName string) {
	time.Sleep(1 * time.Second)
	if err := os.Remove(fileName); err != nil {
		t.Logf("failed to remove the sample data file: %v", err.Error())
	}
}
