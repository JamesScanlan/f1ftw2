package json_converter

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type OriginalGrandPrixRaceResult struct {
	Driver     string `json:"name"`
	Team       string `json:"team"`
	Qualifying string `json:"qualifying"`
	Sprint     string `json:"sprint"`
	Grid       string `json:"grid"`
	Position   string `json:"position"`
}

type OriginalGrandPrixRace struct {
	Name    string                        `json:"Grand_Prix"`
	Results []OriginalGrandPrixRaceResult `json:"Results"`
}

type OriginalGrandPrixSeason struct {
	Year  string                  `json:"Year"`
	Races []OriginalGrandPrixRace `json:"Races"`
}

type OriginalGrandPrixSeasons struct {
	GrandPrixSeasons []OriginalGrandPrixSeason `json:"Grands_Prix"`
}

func ReadGPData(filename string) OriginalGrandPrixSeasons {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var seasons OriginalGrandPrixSeasons
	json.Unmarshal(byteValue, &seasons)

	return seasons
}
