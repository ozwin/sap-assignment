package models

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_readTrailsDataFromCSV(t *testing.T) {

	fileName := "sample_data"
	createSampledataFile(t, fileName)
	defer deleteSampleDataFile(t, fileName)

	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    Trails
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Valid Data",
			args: args{fileName: fileName},
			want: Trails{{
				Address:    "Test Address 1",
				AccessName: "Test Access 1",
				Difficulty: "Difficult",
			}, {
				Address:    "Test Address 2",
				AccessName: "Test Access 2",
				Difficulty: "Easy",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readTrailsDataFromCSV(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("readTrailsDataFromCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readTrailsDataFromCSV() = %v, want %v", got, tt.want)
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
	data := `Address,Fee,AccessName,ADAtrail
	Test Address 1,Yes,Test Access 1,Difficult
	Test Address 2,Yes,Test Access 2,Easy`
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
