package json_converter

import (
	"encoding/json"
	"os"
	"strconv"
)

type GrandPrixRaceResult struct {
	Driver     string `json:"name"`
	Team       string `json:"team"`
	Qualifying int    `json:"qualifying"`
	Sprint     int    `json:"sprint"`
	Grid       int    `json:"grid"`
	Position   int    `json:"position"`
}

type GrandPrixRace struct {
	Name    string                `json:"Grand_Prix"`
	Results []GrandPrixRaceResult `json:"Results"`
}

type GrandPrixSeason struct {
	Year  int             `json:"Year"`
	Races []GrandPrixRace `json:"Races"`
}

type GrandPrixSeasons struct {
	GrandPrixSeasons []GrandPrixSeason `json:"Grands_Prix"`
}

func WriteGPData(filename string, convertedGrandPrixSeasons GrandPrixSeasons) {
	file, _ := json.MarshalIndent(convertedGrandPrixSeasons, "", " ")
	_ = os.WriteFile(filename, file, 0644)
}

func ConvertToNewGPData(originalGrandPrixSeasons OriginalGrandPrixSeasons) GrandPrixSeasons {
	var convertedGrandPrixSeasons GrandPrixSeasons

	for i := 0; i < len(originalGrandPrixSeasons.GrandPrixSeasons); i++ {
		originalSeason := originalGrandPrixSeasons.GrandPrixSeasons[i]
		var convertedSeason GrandPrixSeason
		convertedSeason.Year, _ = strconv.Atoi(originalSeason.Year)
		for j := 0; j < len(originalSeason.Races); j++ {
			originalRace := originalSeason.Races[j]
			var convertedRace GrandPrixRace
			convertedRace.Name = originalRace.Name
			for k := 0; k < len(originalRace.Results); k++ {
				originalResult := originalRace.Results[k]
				var convertedResult GrandPrixRaceResult
				convertedResult.Driver = originalResult.Driver
				convertedResult.Team = originalResult.Team
				convertedResult.Qualifying, _ = strconv.Atoi(originalResult.Qualifying)

				parseSprint(originalResult, &convertedResult)
				// convertedResult.Sprint, _ = strconv.Atoi(originalResult.Sprint)
				convertedResult.Grid, _ = strconv.Atoi(originalResult.Grid)
				convertedResult.Position, _ = strconv.Atoi(originalResult.Position)
				convertedRace.Results = append(convertedRace.Results, convertedResult)
			}
			convertedSeason.Races = append(convertedSeason.Races, convertedRace)
		}
		convertedGrandPrixSeasons.GrandPrixSeasons = append(convertedGrandPrixSeasons.GrandPrixSeasons, convertedSeason)
	}
	return convertedGrandPrixSeasons
}

func parseSprint(originalResult OriginalGrandPrixRaceResult, newResult *GrandPrixRaceResult) {
	parsedValue, err := strconv.Atoi(originalResult.Sprint)
	if err == nil {
		if parsedValue != 0 {
			newResult.Sprint = parsedValue
		}
	}
}
