package json_converter

import (
	"encoding/json"
	"os"
	"strconv"
	"time"
)

type NewPrediction struct {
	Predictor   string `json:"Predictor"`
	Qualifying  string `json:"Qualifying"`
	Race        string `json:"Race"`
	Progression string `json:"Progression"`
}

type NewGrandPrix struct {
	Name        string          `json:"Grand_Prix"`
	Predictions []NewPrediction `json:"Predictions"`
}

type NewSeason struct {
	Year  int            `json:"Year"`
	Races []NewGrandPrix `json:"Races"`
}

type NewSeasons struct {
	Seasons []NewSeason `json:"Grands_Prix"`
}

func WritePredictions(filename string, newSeasons NewSeasons) {
	file, _ := json.MarshalIndent(newSeasons, "", " ")
	_ = os.WriteFile(filename, file, 0644)
}

func GetCurrentYearPredictions(seasons NewSeasons) NewSeasons {
	var current_seasons NewSeasons
	for i := 0; i < len(seasons.Seasons); i++ {
		season := seasons.Seasons[i]
		if season.Year == time.Now().Year() {
			current_seasons.Seasons = append(current_seasons.Seasons, season)
		}
	}
	return current_seasons
}

func ConvertToNewPredictions(orginal_seasons OriginalSeasons) NewSeasons {
	var new_seasons NewSeasons
	for i := 0; i < len(orginal_seasons.Seasons); i++ {
		season := orginal_seasons.Seasons[i]
		var new_season NewSeason
		new_season.Year, _ = strconv.Atoi(season.Year)
		for j := 0; j < len(season.Races); j++ {
			race := season.Races[j]
			var new_race NewGrandPrix
			new_race.Name = race.Name
			for k := 0; k < len(race.Predictions); k++ {
				prediction := race.Predictions[k]
				var new_prediction NewPrediction
				new_prediction.Predictor = prediction.Predictor
				new_prediction.Qualifying = prediction.Qualifying
				new_prediction.Race = prediction.Race
				new_prediction.Progression = prediction.Progression
				new_race.Predictions = append(new_race.Predictions, new_prediction)
			}
			new_season.Races = append(new_season.Races, new_race)
		}
		new_seasons.Seasons = append(new_seasons.Seasons, new_season)
	}
	return new_seasons
}
