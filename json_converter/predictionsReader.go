package json_converter

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
)

type OriginalPrediction struct {
	Predictor   string `json:"Predictor"`
	Qualifying  string `json:"Qualifying"`
	Race        string `json:"Race"`
	Progression string `json:"Progression"`
}

type OriginalGrandPrix struct {
	Name        string               `json:"Grand_Prix"`
	Predictions []OriginalPrediction `json:"Predictions"`
}

type OriginalSeason struct {
	Year  string              `json:"Year"`
	Races []OriginalGrandPrix `json:"Races"`
}

type OriginalSeasons struct {
	Seasons []OriginalSeason `json:"Grands_Prix"`
}

// ReadJSONasText reads the JSON data from a file and returns it as a string
func readJSONasText() string {
	return `{"Grands_Prix":[{"Year":"2024","Races":[{"Grand_Prix":"Netherlands","Predictions":[{"Predictor":"Andrew Chadwick","Qualifying":"Lewis Hamilton","Race":"Carlos Sainz Jnr","Progression":"Williams"},{"Predictor":"James Scanlan","Qualifying":"Charles Leclerc","Race":"Max Verstappen","Progression":"Mercedes"},{"Predictor":"Matt Potter","Qualifying":"Alexander Albon","Race":"George Russell","Progression":"Alpine"}]}]}]}`
}

func ListCurrentDirectory() {
	fsys := os.DirFS(".")
	fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	})
}

func ReadPredictionsJSON(filename string) OriginalSeasons {

	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var seasons OriginalSeasons
	json.Unmarshal(byteValue, &seasons)

	return seasons
}
